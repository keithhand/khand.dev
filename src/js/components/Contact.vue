<template>
  <section id="contact">
    <h4><strong>{{ name }}</strong></h4>
    <ul>
      <li>{{ phone }}</li>
      <li>{{ email }}</li>
    </ul>
    <hr>
  </section>
</template>

<script>
export default {
  props: {
    name: {
      type: String,
      required: true,
    },
    phone: {
      type: Number,
      required: true,
    },
    email: {
      type: String,
      required: true,
    },
  },

  data() {
    return {
      hostname: window.location.hostname,
    };
  },
};
</script>
