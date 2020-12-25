'use strict';

const Clients = require('./balancers/client');

const client = new Clients('localhost', 8080);

// Scenario 1: Get info about balancers
(async () => {
    try {
        const balancers = await client.getBalancers()
        console.log('---Scenario 1---')
        balancers.forEach(balancer => {
            console.log(balancer)
        })
    } catch(err) {
        console.log(err)
    }
})();

// Scenario 2: Update existed machine id
(async () => {
    try {
        const result = await client.updateMachines(1, true)
        console.log('---Scenario 2---')
        console.log(result)
    } catch(err) {
        console.log(err)
    }
})();

// Scenario 3: Get info about balancers after update
(async () => {
    try {
        const balancers = await client.getBalancers()
        console.log('---Scenario 3---')
        balancers.forEach(balancer => {
            console.log(balancer)
        })
    } catch(err) {
        console.log(err)
    }
})();

// Scenario 4: Update machine with invalid id
(async () => {
    try {
        const result = await client.updateMachines(-3, true)
        console.log('---Scenario 4---')
        console.log(result)
    } catch(err) {
        console.log(err)
    }
})();
