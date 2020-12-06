'use strict';

const Clients = require('./balancers/client');

const client = new Clients('http://localhost:8080');

// Scenario 1: Get info about balancers


(async () => {
    console.log('---Scenario 1---')
    const transactions = await client.getBalancers()
    transactions.forEach(balancer => {
        console.log(balancer)
    })
})();


// Scenario 2: Update machine state
