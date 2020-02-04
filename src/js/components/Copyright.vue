<template>
  <section id="copyright">
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
