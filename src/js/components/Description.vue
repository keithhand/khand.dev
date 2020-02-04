<template>
  <section id="description">
    <h2>{{ header }}</h2>
    <p>{{ content }}</p>
  </section>
</template>

<script>
export default {
  props: {
    header: {
      type: String,
      required: true,
    },
    content: {
      type: String,
      required: true,
    },
  },
};
</script>
