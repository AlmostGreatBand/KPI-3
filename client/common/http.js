'use strict';

const http = require('http');

class Client {
    static #METHODS = {
        GET: 'GET',
        PUT: 'PUT',
        POST: 'POST',
    }

    constructor(baseUrl, port) {
        this.baseUrl = baseUrl;
        this.port = port;
    }

    async get(path) {
        const params = this.#createParams(path, Client.#METHODS.GET);
        return await this.#request(params);
    }

    async put(path, data) {
        const params = this.#createParams(path, Client.#METHODS.PUT);
        return await this.#request(params, JSON.stringify(data));
    }

    #createParams = (path, method) => ({
        hostname: this.baseUrl,
        port: this.port,
        path: path,
        method: method,
        headers: {
            'Content-Type': 'application/json; charset=UTF-8',
        },
    });

    #request = async (params, body) => new Promise((resolve, reject) => {
        const req = http.request(params);
        req.on('response', res => {
            const code = res.statusCode.toString();
            if (code === '204') {
                /*
                    if in function we have only 'end' subscriber it won't be emitted
                    so I add empty on end subscriber
                */
                res.on('data', () => {})

                res.on('end', () => {
                    resolve(`Response code: ${code} - machine updated`)
                })
            } else if (code.match(/^2\d\d$/)) {
                let stream = '';
                res.on('data', chunk => {
                    stream += chunk;
                });

                res.on('end', () => {
                    try {
                        resolve(JSON.parse(stream));
                    } catch(e) {
                        const errMessage = `Error: something went wrong while trying to parse http data: ${e}`;
                        reject(errMessage)
                    }
                })
            } else {
                const errMessage = `Error: code ${code} message: ${res.statusMessage}`;
                reject(errMessage)
            }
        });

        req.on('error', err => {
            const errMessage = `Error: code ${err.code} message: ${err.message}`;
            reject(errMessage)
        });

        if (body) {
            req.write(body);
        }
        req.end();
    });
}

module.exports = Client;