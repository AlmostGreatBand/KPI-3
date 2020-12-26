'use strict';

const Clients = require('./balancers/client');

const client = new Clients('localhost', 8080);

const getBalancersInfo = async () => {
    try {
        const balancers = await client.getBalancers()
        balancers.forEach(balancer => {
            console.log(balancer)
        })
    } catch(err) {
        console.log(err)
    }
}

const updateMachines = async (id, state) => {
    try {
        const result = await client.updateMachines(id, state)
        console.log(result)
    } catch(err) {
        console.log(err)
    }
}

// Usage:

(async () => {
    // Scenario 1: Get info about balancers
    console.log('---Scenario 1---')
    await getBalancersInfo()

    // Scenario 2: Update existed machine id
    console.log('---Scenario 2---')
    await updateMachines(1, false)

    // Scenario 3: Get info about balancers after update
    console.log('---Scenario 3---')
    await getBalancersInfo()

    // Scenario 4: Update machine with invalid id
    console.log('---Scenario 4---')
    await updateMachines(-3, true)
})();
