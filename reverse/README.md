Lenses! 

In this directory we have a cue file with:
- three schemas: 0.0, 0.1, and 1.0
- two lenses, for 0.1->1.0 and 1.0->0.1

Translate from both 0.0 and 0.1 to 1.0 work as expected - the minor version lens is implied.
The reverse translation panics, from 1.0 to either 0.1 or 0.1

How many lenses is thema expecting in this example?