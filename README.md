# Little Jake

Little Jake is a command line interface to use Arcehage web API in Discord.

This only works on Korea servers yet.

## function list

implemented
- Route bond region
- Load user's information

Unimplemented
- Keep an eyes on user's changes
- Check price of item in game
- Calculate the time until growth of crops and alarm it
- Set an alarm the time to go to the raid
- Auction raid item

## Usage

###[prefix][invoke] [arg1] [arg2] [arg3]...

prefix : '!'

invokes :
- 'h' or 'help'
- 'b' or 'bond'
- 'u' or 'user'
- 'f' or 'follow'
- 'p' or 'price'
- 'c' or 'crops'
- 'r' or 'raid'
- 'a' or 'auction'

args : 
- -flag value
- value

### Example
- ```!user DevMyong or !user DevMyong@NUI```
- ```!bond or !bond NUI```
- ```!price ITEM_NAME```
- ```!crops```
- ```!raid RedDragon|Kadoom|War```

