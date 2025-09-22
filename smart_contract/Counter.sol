// SPDX-License-Identifier: MIT

pragma solidity ^0.8.20;


contract Counter{
    uint256 public counter;

    event CounterIncremented(uint256 newCounter);
    event CounterDecremented(uint256 newCounter);
    event CounterReset(uint256 newCounter);

    function incrementCounter() public {
        if (counter >= 1000000) {
            reset();
        }
        counter++;

        emit CounterIncremented(counter);
    }

    function getCounter() public view returns (uint256) {
        require(counter <= 1000000, "Counter is too large");
        require(counter >= 0, "Counter is too small");
        return counter;
    }

    function decrementedCounter() public {
        require(counter > 0, "Counter is too small");
        counter--;
        emit CounterDecremented(counter);
    }

    function reset() internal  {
        counter = 0;
        emit CounterReset(counter);
    }
}