<template>
  <section id="projects">
    <h2>My Projects</h2>
    <hr>
    <ul>
      <li
        v-for="p in projects"
        :key="p.title"
      >
        <p>
          {{ p.name }}
        </p>
      </li>
    </ul>
  </section>
</template>

<script>
export default {
  props: {
    projects: {
      type: Array,
      required: true,
    },
  },
};
</script>
