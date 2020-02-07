import Vue from 'vue';

// fontawesome imports
import { library } from '@fortawesome/fontawesome-svg-core';
import { FontAwesomeIcon, FontAwesomeLayers } from '@fortawesome/vue-fontawesome';
// fas icons
import {
  faCircle, faTerminal, faDatabase, faLaptopCode, faEnvelope, faPhone,
} from '@fortawesome/free-solid-svg-icons';
// fab icons
import {
  faHtml5, faPhp, faJsSquare, faCss3Alt, faPython, faJava, faGit, faDocker, faLinux, faWindows,
  faApple, faGithub,
} from '@fortawesome/free-brands-svg-icons';
// far icons
import {
  faFileAlt, faGem,
} from '@fortawesome/free-regular-svg-icons';
import App from './App.vue';
// add all fontawesome icons to library object
library.add(faHtml5, faPhp, faJsSquare, faCss3Alt, faPython, faJava, faGit, faDocker, faLinux,
  faWindows, faApple, faCircle, faTerminal, faDatabase, faFileAlt, faGem, faLaptopCode,
  faEnvelope, faPhone, faGithub);

Vue.component('font-awesome-icon', FontAwesomeIcon);
Vue.component('font-awesome-layers', FontAwesomeLayers);

Vue.config.productionTip = false;

new Vue({ // eslint-disable-line no-new
  el: '#app',
  render: (h) => h(App),
});
