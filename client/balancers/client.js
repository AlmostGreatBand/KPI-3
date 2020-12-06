'use strict';

const Client = require('../common/http');

class BalancerClient extends Client {
    constructor(baseUrl, port) {
        super(baseUrl, port);
    }

    async getBalancers() {
        return await this.get('/balancers')
    }

    async updateMachines(id, state) {
        return await this.put('/balancers', { id: id, state: state })
    }
}

module.exports = BalancerClient;