### Requirements

* Go 1.15 Or higher

### Framework

This project utilizes Gorilla web toolkit.

### Usage

1. Hello Endpoint - /hello

   This endpoint greets people, in this case wizeline.
   ```
   Eg. http://localhost:8000/hello
   ```

2. Card Game Endpoint - /play-cards

   This endpoint draw two cards (values and suits) from http://deckofcardsapi.com/, one for the computer and one for the user, then it compare the value of the cards, being the values:

   ACE - 14
   KING - 13
   QUEEN - 12
   JACK - 11
   All the other values correspond to the number of the card.

   Then it decides if the user won, lose or if it was a tie.

   ```
   Eg. http://localhost:8000/play-cards
   ```