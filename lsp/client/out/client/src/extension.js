"use strict";
/* --------------------------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Licensed under the MIT License. See License.txt in the project root for license information.
 * ------------------------------------------------------------------------------------------ */
Object.defineProperty(exports, "__esModule", { value: true });
exports.deactivate = exports.activate = void 0;
const net = require("net");
const main_1 = require("../../vendor/vscode-languageserver-node/client/src/node/main");
// import {
// 	LanguageClient,
// 	LanguageClientOptions,
// 	ServerOptions,
// 	TransportKind,
//     StreamInfo,
// } from 'vscode-languageclient';
let client;
function activate(context) {
    console.log("configing client...");
    const disposable = startLanguageServerTCP(5007, ["plaintext"]);
    context.subscriptions.push(disposable);
}
exports.activate = activate;
function startLanguageServer(command, args, documentSelector) {
    const serverOptions = {
        command,
        args,
    };
    const clientOptions = {
        documentSelector: documentSelector,
        synchronize: {
            configurationSection: "plaintext"
        },
    };
    return new main_1.LanguageClient(command, serverOptions, clientOptions).start();
}
function startLanguageServerTCP(address, documentSelector) {
    const serverOptions = () => {
        return new Promise((resolve, reject) => {
            const client = new net.Socket();
            client.connect(address, "127.0.0.1", () => {
                resolve({ reader: client, writer: client });
            });
        });
    };
    const clientOptions = {
        documentSelector: documentSelector,
    };
    const client = new main_1.LanguageClient(`tcp language server (port ${address})`, serverOptions, clientOptions);
    const disposable = client.start();
    console.log("client has started the tcp simple socket");
    client.onDidChangeState(e => {
        console.log("did change state");
        console.log(e);
    });
    // client.initializeResult.then((result) => {
    //     console.log("client initialized");
    //     console.log(result);
    // });
    console.log("about to return disposable...");
    return disposable;
}
function deactivate() {
    if (!client) {
        return undefined;
    }
    return client.stop();
}
exports.deactivate = deactivate;
