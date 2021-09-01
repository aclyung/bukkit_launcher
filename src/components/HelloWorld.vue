<template>
  <v-container>
    <body>
    <v-row class="text-center">
      <v-col cols="12">
        <v-img
          :src="require('../assets/getbukkit.png')"
          class="my-3"
          contain
          height="200"
        />
      </v-col>

      <v-col class="mb-4">
        <h1 class="display-2 font-weight-bold mb-3">
          WELCOME TO <br>BUKKIT LAUNCHER
        </h1>
        <v-btn @click="hi">hi</v-btn>
      </v-col>

    </v-row>
    </body>
  </v-container>
</template>

<script>
// eslint-disable-next-line import/no-duplicates
import { Terminal } from 'xterm';
import { FitAddon } from 'xterm-addon-fit';
import 'xterm/css/xterm.css';
// eslint-disable-next-line import/no-duplicates
import 'xterm/lib/xterm';
import './wasm_exec';

// eslint-disable-next-line import/no-extraneous-dependencies
const { dialog } = require('electron').remote;

const fs = require('fs');

if (!WebAssembly.instantiateStreaming) {
  WebAssembly.instantiateStreaming = async (resp, importObject) => {
    const source = await (await resp).arrayBuffer();
    // eslint-disable-next-line no-return-await
    return await WebAssembly.instantiate(source, importObject);
  };
}
// eslint-disable-next-line no-undef
const go = new Go();

let mod;
let inst;
WebAssembly.instantiateStreaming(fetch('./getbukkit.wasm'), go.importObject).then((result) => {
  mod = result.module;
  inst = result.instance;
}).catch((err) => {
  console.error(err);
});

// eslint-disable-next-line no-unused-vars
async function run() {
  console.clear();
  await go.run(inst);
  inst = await WebAssembly.instantiate(mod, go.importObject); // reset instance
}

export default {
  name: 'HelloWorld',
  data: () => ({
    term: '',
    socket: '',
  }),
  mounted() {
    const url = 'ws://***********';
    this.init(url);
  },
  methods: {
    hi() {
      dialog.showOpenDialogSync();
      fs.writeFileSync('./hi.txt', 'helllo');
    },
    initXterm() {
      this.term = new Terminal({
        rendererType: 'canvas',
        rows: 35,
        convertEol: true,
        scrollback: 10,
        disableStdin: false,
        cursorStyle: 'underline',
        cursorBlink: true,
        theme: {
          foreground: 'yellow',
          background: '#060101',
          cursor: 'help',
        },
      });
      this.term.open(document.getElementById('terminal'));
      const fitAddon = new FitAddon();
      this.term.loadAddon(fitAddon);
      //  Support input and paste methods
      const tthis = this;
      this.term.onData((key) => {
        const order = ['stdin', key];
        tthis.socket.onsend(JSON.stringify(order));
      });
    },
    init(url) {
      //  Instantiate socket
      this.socket = new WebSocket(url);
      //  Monitor socket connection
      this.socket.onopen = this.open;
      //  Monitor socket error messages
      this.socket.onerror = this.error;
      //  Listen for socket messages
      this.socket.onmessage = this.getMessage;
      //  Send socket message
      this.socket.onsend = this.send;
    },
    open() {
      console.log('Socket connection is successful');
      this.initXterm();
    },
    error() {
      console.log('error in connecting');
    },
    close() {
      this.socket.close();
      console.log('socket is closed');
    },
    getMessage(msg) {
      this.term.write(JSON.parse(msg.data)[1]);
    },
    send(order) {
      this.socket.send(order);
    },
  },
};
</script>

<style>
.box {
  width: 100%;
  height: 100%;
}
</style>
