<template>
  <div>
    <LoginForm class="form"/>
  </div>
</template>

<script>
import LoginForm from '@/components/LoginForm.vue';

export default {
  data: function () {
    return {
    };
  },
  components: {
    LoginForm,
  },
}
</script>

<style scoped>
  @media only screen and (min-width: 768px) {
    .form {
      max-width: 50%;
    }
  }
</style>