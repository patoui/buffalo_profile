require("expose-loader?$!expose-loader?jQuery!jquery");
require("bootstrap/dist/js/bootstrap.bundle.js");
require("@fortawesome/fontawesome-free/js/all.js");

import Vue from 'vue/dist/vue.js';

 Vue.component('tags', require('./components/Tags.vue').default);

 const app = new Vue({ el: '#app' });
