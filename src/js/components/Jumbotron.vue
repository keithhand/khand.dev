<template>
  <section id="jumbotron">
    <h1>{{ name }}</h1>
    <h3>{{ title }}</h3>
    <a>Who am I?</a>
  </section>
</template>

<script>
export default {
  props: {
    name: {
      type: String,
      required: true,
    },
    title: {
      type: String,
      required: true,
    },
  },
};
</script>
