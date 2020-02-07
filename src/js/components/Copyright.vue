<template>
  <section id="copyright">
    <hr>
    <p>Copyright © {{ new Date().getFullYear() }} {{ hostname }}</p>
  </section>
</template>

<script>
export default {
  data() {
    return {
      hostname: window.location.hostname,
    };
  },
};
</script>

<style lang="scss" scoped>
  @import "../../assets/css/variables";

  section {
    background: $main-color;
    color: white;
    font-size: 14px;
    padding: 0 0 10px 0;
    margin: -20px 0 0;

    hr {
      margin-bottom: 20px;
    }

    p {
      margin: 0;
    }
  }
</style>
