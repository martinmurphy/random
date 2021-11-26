# Random

Choose a set of random items from a list

## Usage

If a list of playing card names are stored in a file `cards.txt`, choose 5 at ramdom
 

```
./choose_items --file cards.txt -c 5
```

```
Choosing 5 items from 52 in file: cards.txt
place: 1, 'J of Clubs' (37)
place: 2, 'Q of Hearts' (12)
place: 3, '4 of Diamonds' (17)
place: 4, 'Q of Clubs' (38)
place: 5, '6 of Spades' (45)
```

Output only the names of the randomly chosen items (no position information)
```
./choose_items --file cards.txt -c 5 --onlyChoices
```

```
A of Spades
9 of Spades
8 of Hearts
6 of Hearts
6 of Diamonds
```

Specify the total number of items in the input file, to output it in a random (shuffled) order


```
./choose_items --file cards.txt -c 52 --onlyChoices
```

## Build

```
go build choose_items.go choose.go
```