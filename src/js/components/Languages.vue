<template>
  <section id="languages">
    <h2>My Knowledge Base</h2>
    <hr>
    <ul>
      <li
        v-for="l in languages"
        :key="l.name"
      >
        {{ l.title }}
      </li>
    </ul>
  </section>
</template>

<script>
export default {
  props: {
    languages: {
      type: Array,
      required: true,
    },
  },
};
</script>
