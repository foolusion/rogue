# rogue

## main loop

1. wait for input from player
   1.1. if it is a direction input check if it is attack or move
   1.2. if attack run attack logic
   1.3. if move check, check if it is legal
      1.3.1. if it is not legal wait for new input
      1.3.2. if it is a bad move warn player and wait for new input
      1.3.3. if it is good move do move
2. update player
3. update everyone else
4. draw screen
5. go to 1
