/* --------------------------------------------------------------------------------------------
 * Copyright (c) Microsoft Corporation. All rights reserved.
 * Licensed under the MIT License. See License.txt in the project root for license information.
 * ------------------------------------------------------------------------------------------ */

import * as path from 'path';
import { workspace, ExtensionContext, TextDocument, Disposable } from 'vscode';
import * as net from 'net';
import {
	LanguageClient,
	LanguageClientOptions,
	ServerOptions,
} from 'vscode-languageclient';

let client: LanguageClient;

export function activate(context: ExtensionContext) {
    console.log("configuring workspace...");
    context.subscriptions.push(startLanguageServerTCP(5007, ["plaintext"]));
}

function startLanguageServerTCP(address: number, documentSelector: string[]): Disposable {
    const serverOptions: ServerOptions = () => {
        return new Promise((resolve, reject) => {
            const client = new net.Socket();
            client.connect(address, "127.0.0.1", () => {
                resolve({reader: client, writer: client});
            })
        })
    }

    const clientOptions: LanguageClientOptions = {
        documentSelector: documentSelector,
    }


    const client = new LanguageClient(`tcp language server (port ${address})`, serverOptions, clientOptions)
    const disposable = client.start();

    return disposable;
}

export function deactivate(): Thenable<void> | undefined {
	if (!client) {
		return undefined;
	}
	return client.stop();
}





