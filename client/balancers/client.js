'use strict';

const Client = require('../common/http');

class BalancerClient extends Client {
    constructor(baseUrl) {
        super(baseUrl);
    }

    getBalancers() {
        return this.get('/balancers')
    }

    updateMachines(id, state) {
        return this.put('/balancers', { id: id, state: state })
    }
}

module.exports = BalancerClient;
