<template>
  <section id="description">
    <h2>{{ header }}</h2>
    <p>{{ content }}</p>
  </section>
</template>

<script scoped>
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

<style lang="scss" scoped>
  @import "../../assets/css/variables";

  section {
    background: white;
    color: $secondary-color;
    margin: -50px 0 0 0;
    padding: 0 95px 60px 95px;
  }
</style>
