'use strict';

const http = require('http');

class Client {
    static #METHODS = {
        GET: 'GET',
        PUT: 'PUT',
        POST: 'POST',
    }

    constructor(baseUrl) {
        this.baseUrl = baseUrl;
    }

    async get(path) {
        const params = this.#createParams(path, Client.#METHODS.GET);
        try {
            return this.#request(params);
        } catch (err) {
            console.log(err);
        }
    }

    async put(path, data) {
        const params = this.#createParams(path, Client.#METHODS.PUT);
        try {
            return this.#request(params,JSON.stringify(data));
        } catch (err) {
            console.log(err);
        }
    }

    #createParams = (path, method) => ({
        hostname: this.baseUrl,
        path: path,
        method: method,
        headers: {
            'Content-Type': 'application/json; charset=UTF-8',
        },
    });

    #request = async (params, body) => {
        const req = http.request(params);

        req.on('response', res => {
            const code = res.statusCode.toString();
            if (code.match(/^2\d\d$/)) {
                let stream = '';

                res.on('data', chunk => {
                    stream += chunk;
                });

                res.on('end', () => {
                    try {
                        return JSON.parse(stream);
                    } catch(e) {
                        const errMessage = `Error: something went wrong while trying to parse http data: ${e}`;
                        throw new Error(errMessage);
                    }
                })
            } else {
                const errMessage = `Error: code ${code} message: ${res.statusMessage}`;
                throw new Error(errMessage);
            }
        });

        req.on('error', err => {
            const errMessage = `Error: code ${err.code} message: ${err.message}`;
            throw new Error(errMessage);
        });

        if (body) {
            req.write(body);
        }
        req.end();
    }
}

module.exports = Client;
