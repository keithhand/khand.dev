<template>
  <div id="app">
    <header>
      <jumbotron
        :name="contact.name"
        :title="contact.title"
      />
    </header>
    <main>
      <aside>
        <description
          :header="description.header"
          :content="description.content"
        />
      </aside>
      <languages
        :languages="languages"
      />
      <projects
        :projects="project_list"
      />
      <contact
        :name="contact.name"
        :phone="contact.phone"
        :email="contact.email"
      />
    </main>
    <footer>
      <copyright />
    </footer>
  </div>
</template>

<script>
import _ from 'lodash';

import Jumbotron from './js/components/Jumbotron.vue';
import Description from './js/components/Description.vue';
import Languages from './js/components/Languages.vue';
import Projects from './js/components/Projects.vue';
import Contact from './js/components/Contact.vue';
import Copyright from './js/components/Copyright.vue';

import languages from './assets/languages.json';

export default {
  components: {
    Jumbotron,
    Description,
    Languages,
    Projects,
    Contact,
    Copyright,
  },

  data() {
    return {
      contact: {
        name: 'Keith Hand',
        title: 'DevOps Engineer / Software Developer',
        email: 'keith@hand.technology',
        phone: '(405) 574 5813',
      },
      description: {
        header: 'Hello, my name is Keith, and I like to write code!',
        content: 'I enjoy taking on any project, big or small. Whether you need a personal webpage, you have the next big AAA iOS title, or you need someone to manage your webserver, I want to be the guy to help make your dreams become a reality.',
      },
    };
  },

  computed: {
    languages: () => languages,

    project_list() {
      return _.range(4).map((i) => ({
        name: `project-${i + 1}`,
        image_url: 'https://source.unsplash.com/random/800x600',
      }));
    },
  },
};
</script>

<style lang="scss">
  @import "assets/css/variables";

  body {
    margin: 0;
  }

  #app {
    font-family: 'Avenir', Helvetica, Arial, sans-serif;
    -webkit-font-smoothing: antialiased;
    -moz-osx-font-smoothing: grayscale;
    text-align: center;
    color: #2c3e50;
  }

  h1, h2 {
    font-weight: normal;
  }

  ul {
    list-style-type: none;
    padding: 0;
  }

  li {
    display: inline-block;
    margin: 0 10px;
  }

  a {
    color: #42b983;
  }

  hr {
    width: 100px;
  }

  section {
    padding: 50px 0;
  }
</style>
