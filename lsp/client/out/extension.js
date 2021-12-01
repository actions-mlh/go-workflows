"use strict";
/* --------------------------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Licensed under the MIT License. See License.txt in the project root for license information.
 * ------------------------------------------------------------------------------------------ */
Object.defineProperty(exports, "__esModule", { value: true });
exports.deactivate = exports.activate = void 0;
const net = require("net");
const vscode_languageclient_1 = require("vscode-languageclient");
let client;
function activate(context) {
    console.log("configuring workspace...");
    context.subscriptions.push(startLanguageServerTCP(5007, ["yaml"]));
}
exports.activate = activate;
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
    const client = new vscode_languageclient_1.LanguageClient(`tcp language server (port ${address})`, serverOptions, clientOptions);
    const disposable = client.start();
    return disposable;
}
function deactivate() {
    if (!client) {
        return undefined;
    }
    return client.stop();
}
exports.deactivate = deactivate;
