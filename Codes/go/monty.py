
from tqdm import tqdm
from random import randint

def monty_hall():
    my_door=randint(1,3)
    prize_door=randint(1,3)
    goat1=True
    goat2=True
    while goat1:
        goat_door1=randint(1,3)
        if goat_door1!=prize_door:
            goat1=False
    while goat2:
        goat_door2=randint(1,3)
        if (goat_door2!=prize_door and goat_door2!=goat_door1):
            goat2=False
    switch=True
    show_goat=True
    while show_goat:
        monty_choice=randint(1,3)
        if (monty_choice!=prize_door and monty_choice!=my_door):
            show_goat=False
    while switch:
        new_door=randint(1,3)
        if (new_door!=my_door and new_door!=monty_choice):
            switch=False
    return my_door,new_door,prize_door,goat_door1,goat_door2,monty_choice

def play_game(num_games):
    success=0
    for i in tqdm(range(num_games)):
        my_door,new_door,prize_door,goat_door1,goat_door2,monty_choice=monty_hall()
        if new_door==prize_door:
            success+=1
    prob=float(success)/num_games
    theo_value=2/3.
    error=((abs(prob-theo_value))/theo_value)*100
    print("Theoretical value = %.5f , Monte Carlo simulation value = %s, error = %s %%"%(theo_value,prob,error))

def main():
    num_games=int(input('Type the number of games played in the simulation: '))
    play_game(num_games)

if __name__ == '__main__':
    main()