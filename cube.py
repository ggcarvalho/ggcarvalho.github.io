import numpy as np
import matplotlib.pyplot as plt
from mpl_toolkits import mplot3d

fig = plt.figure()
ax = plt.axes(projection="3d")

t = np.linspace(0, 1, 1000, endpoint=True)
for i in t:
    ax.scatter3D(i, i, i, color=(i,i,i))
    ax.scatter3D(i, 0, 0, color=(i,0,0))
    ax.scatter3D(0, i, 0, color=(0,i,0))
    ax.scatter3D(0, 0, i, color=(0,0,i))
plt.show()