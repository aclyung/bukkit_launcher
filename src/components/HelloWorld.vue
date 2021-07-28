<template>
  <v-container>
    <v-row class="text-center">
      <v-col cols="12">
        <v-img
          :src="require('../assets/logo.svg')"
          class="my-3"
          contain
          height="200"
        />
      </v-col>

      <v-col class="mb-4">
        <h1 class="display-2 font-weight-bold mb-3">
          Welcome to Vuetify
        </h1>

        <p class="subheading font-weight-regular">
          For help and collaboration with other Vuetify developers,
          <br>please join our online
          <a
            href="https://community.vuetifyjs.com"
            target="_blank"
          >Discord Community</a>
        </p>
        <v-btn @click="hi">hi</v-btn>
      </v-col>

      <v-col
        class="mb-5"
        cols="12"
      >
        <h2 class="headline font-weight-bold mb-3">
          What's next?
        </h2>
        <div class="box">
          <div id="terminal"></div>
        </div>
        <v-row justify="center">
          <a
            v-for="(next, i) in whatsNext"
            :key="i"
            :href="next.href"
            class="subheading mx-3"
            target="_blank"
          >
            {{ next.text }}
          </a>
        </v-row>
      </v-col>

      <v-col
        class="mb-5"
        cols="12"
      >
        <h2 class="headline font-weight-bold mb-3">
          Important Links
        </h2>

        <v-row justify="center">
          <a
            v-for="(link, i) in importantLinks"
            :key="i"
            :href="link.href"
            class="subheading mx-3"
            target="_blank"
          >
            {{ link.text }}
          </a>
        </v-row>
      </v-col>

      <v-col
        class="mb-5"
        cols="12"
      >
        <h2 class="headline font-weight-bold mb-3">
          Ecosystem
        </h2>

        <v-row justify="center">
          <a
            v-for="(eco, i) in ecosystem"
            :key="i"
            :href="eco.href"
            class="subheading mx-3"
            target="_blank"
          >
            {{ eco.text }}
          </a>
        </v-row>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
// eslint-disable-next-line import/no-duplicates
import { Terminal } from 'xterm';
import { FitAddon } from 'xterm-addon-fit';
import 'xterm/css/xterm.css';
// eslint-disable-next-line import/no-duplicates
import 'xterm/lib/xterm';

const fs = require('fs');

export default {
  name: 'HelloWorld',
  data: () => ({
    term: '',
    socket: '',
    ecosystem: [
      {
        text: 'vuetify-loader',
        href: 'https://github.com/vuetifyjs/vuetify-loader',
      },
      {
        text: 'github',
        href: 'https://github.com/vuetifyjs/vuetify',
      },
      {
        text: 'awesome-vuetify',
        href: 'https://github.com/vuetifyjs/awesome-vuetify',
      },
    ],
    importantLinks: [
      {
        text: 'Documentation',
        href: 'https://vuetifyjs.com',
      },
      {
        text: 'Chat',
        href: 'https://community.vuetifyjs.com',
      },
      {
        text: 'Made with Vuetify',
        href: 'https://madewithvuejs.com/vuetify',
      },
      {
        text: 'Twitter',
        href: 'https://twitter.com/vuetifyjs',
      },
      {
        text: 'Articles',
        href: 'https://medium.com/vuetify',
      },
    ],
    whatsNext: [
      {
        text: 'Explore components',
        href: 'https://vuetifyjs.com/components/api-explorer',
      },
      {
        text: 'Select a layout',
        href: 'https://vuetifyjs.com/getting-started/pre-made-layouts',
      },
      {
        text: 'Frequently Asked Questions',
        href: 'https://vuetifyjs.com/getting-started/frequently-asked-questions',
      },
    ],
  }),
  mounted() {
    const url = 'ws://***********';
    this.init(url);
  },
  methods: {
    hi() {
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
