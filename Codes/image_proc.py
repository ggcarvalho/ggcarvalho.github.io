import imageio
import numpy as np
from tqdm import tqdm

def get_shape(image):
    shape = image.shape
    if len(shape)==3 or len(shape)==2:
        return shape
    else:
        raise Exception("Sorry, something is wrong!")

def is_grayscale(image):
    if len(get_shape(image))==3:
        return False
    elif len(get_shape(image))==2:
        return True
    else:
        raise Exception("Sorry, something is wrong!")

def get_luminance(image):
    return 0.299*image[:, :, 0] + 0.587*image[:, :, 1] + 0.114*image[:, :, 2]

def zeros(height, width, depth=None):
    return np.zeros((height, width)) if depth is None else np.zeros((height, width, depth))

def convert_grayscale(image):
    if not is_grayscale(image):
        height, width, _ = get_shape(image)
        gray_image       = zeros(height, width)
        gray_image = get_luminance(image)
        return gray_image
    else:
        return image

def get_image_range(image):
   return np.min(image), np.max(image)

def adjust_gray(image, new_min, new_max):
   image_min, image_max = get_image_range(image)
   h, w  = get_shape(image)
   adjusted = zeros(h, w)
   adjusted = (image - image_min)*((new_max - new_min)/(image_max - image_min)) + new_min
   return adjusted.astype(np.uint8)

def gen_halftone_masks():
   m = zeros(3, 3, 10)

   m[:, :, 1] = m[:, :, 0]
   m[0, 1, 1] = 1

   m[:, :, 2] = m[:, :, 1]
   m[2, 2, 2] = 1

   m[:, :, 3] = m[:, :, 2]
   m[0, 0, 3] = 1

   m[:, :, 4] = m[:, :, 3]
   m[2, 0, 4] = 1

   m[:, :, 5] = m[:, :, 4]
   m[0, 2, 5] = 1

   m[:, :, 6] = m[:, :, 5]
   m[1, 2, 6] = 1

   m[:, :, 7] = m[:, :, 6]
   m[2, 1, 7] = 1

   m[:, :, 8] = m[:, :, 7]
   m[1, 0, 8] = 1

   m[:, :, 9] = m[:, :, 8]
   m[1, 1, 9] = 1

   return m

def halftone(image):
   gray      = convert_grayscale(image)
   adjusted  = adjust_gray(gray, 0, 9)
   m         = gen_halftone_masks()

   height, width = get_shape(image)
   halftoned        = zeros(3*height, 3*width)
   for j in tqdm(range(height), desc = "halftone"):
       for i in range(width):
           index = adjusted[j, i]
           halftoned[3*j:3+3*j, 3*i:3+3*i] = m[:, :, index]

   halftoned = 255*halftoned
   return halftoned

def clip(a):
    return np.clip(a, 0, 255)

kernels = {"mean"      : np.array([[1/9, 1/9, 1/9],
                                   [1/9, 1/9, 1/9],
                                   [1/9, 1/9, 1/9]]),

           "gaussian"  : np.array([[1/16, 2/16, 1/16],
                                   [2/16, 4/16, 2/16],
                                   [1/16, 2/16, 1/16]]),

           "sharpen"   : np.array([[0 , -1,  0],
                                   [-1,  5, -1],
                                   [0 , -1,  0]]),

           "laplacian" : np.array([[-1, -1, -1],
                                   [-1,  8, -1],
                                   [-1, -1, -1]]),

           "emboss"    : np.array([[-2, -1, 0],
                                   [-1,  1, 1],
                                   [ 0,  1, 2]]),

           "motion"    : np.array([[1/9, 0, 0, 0, 0, 0, 0, 0, 0],
                                   [0, 1/9, 0, 0, 0, 0, 0, 0, 0],
                                   [0, 0, 1/9, 0, 0, 0, 0, 0, 0],
                                   [0, 0, 0, 1/9, 0, 0, 0, 0, 0],
                                   [0, 0, 0, 0, 1/9, 0, 0, 0, 0],
                                   [0, 0, 0, 0, 0, 1/9, 0, 0, 0],
                                   [0, 0, 0, 0, 0, 0, 1/9, 0, 0],
                                   [0, 0, 0, 0, 0, 0, 0, 1/9, 0],
                                   [0, 0, 0, 0, 0, 0, 0, 0, 1/9]]),

           "y_edge"    : np.array([[1 ,  2, 1],
                                   [0 ,  0, 0],
                                   [-1, -2,-1]]),

           "x_edge"    : np.array([[1, 0, -1],
                                   [2, 0, -2],
                                   [1, 0, -1]]),

            "identity" : np.array([[0, 0, 0],
                                   [0, 1, 0],
                                   [0, 0, 0]])}

def apply_kernel(image, kernel):
    kernel_matrix = kernels.get(kernel)
    dim           = len(kernel_matrix)
    center        = (dim - 1)//2

    shape = get_shape(image)
    height, width = shape[0], shape[1]

    if not is_grayscale(image):
        picture = zeros(height, width, 3)

        for y in tqdm(range(height), desc = kernel):
            for x in range(width):

                red = zeros(dim, dim)
                for i in range(dim):
                    for j in range(dim):
                        red[i , j] = image[ (y - center + j)%height, (x - center + i)%width, 0]

                green = zeros(dim, dim)
                for i in range(dim):
                    for j in range(dim):
                        green[i , j] = image[ (y - center + j)%height, (x - center + i)%width, 1]

                blue = zeros(dim, dim)
                for i in range(dim):
                    for j in range(dim):
                        blue[i , j] = image[ (y - center + j)%height, (x - center + i)%width, 2]

                redc   = np.sum(red*kernel_matrix)
                greenc = np.sum(green*kernel_matrix)
                bluec  = np.sum(blue*kernel_matrix)

                r, g, b = map(int,  [redc, greenc, bluec])
                r, g, b = map(clip, [r, g, b])

                picture[y, x, 0] = r
                picture[y, x, 1] = g
                picture[y, x, 2] = b
        return picture
    else:
        picture = zeros(height, width)
        for y in tqdm(range(height), desc = kernel):
            for x in range(width):

                aux = zeros(dim, dim)
                for i in range(dim):
                    for j in range(dim):
                        aux[i , j] = image[ (y - center + j)%height, (x - center + i)%width]
                gray = np.sum(aux*kernel_matrix)

                pxl_intensity = round(gray)
                pxl_intensity = clip(pxl_intensity)
                picture[y, x] = int(pxl_intensity)
        return picture

def transpose(m):
    height, width, depth = get_shape(m)

    transposed = zeros(width, height, depth)
    for i in range(width):
        for j in range(height):
            transposed[i, j] = m[j, i]
    return transposed

def aux90(image):
    return transpose(image)[:,::-1]

def rot90(image):
    print("Rotating the image 90 degrees clockwise...")
    rot = aux90(image)
    return transpose(image)[:,::-1]

def rot180(image):
    print("Rotating the image 180 degrees...")
    rot = image[::-1, ::-1]
    return rot

def rotm90(image):
    print("Rotating the image 90 degrees counterclockwise...")
    rot = aux90(image[::-1, ::-1])
    return rot

def vert_flip(image):
    print("Flipping vertically...")
    flip = image[::-1]
    return flip

def hor_flip(image):
    print("Flipping horizontally...")
    flip = image[:, ::-1]
    return flip

geometric_transforms = {"rot90"     : rot90,
                        "rot180"    : rot180,
                        "rotm90"    : rotm90,
                        "vert_flip" : vert_flip,
                        "hor_flip"  : hor_flip}

def intensity(image, factor):
   return clip(factor*image)

def negative(image):
  return 255 - image