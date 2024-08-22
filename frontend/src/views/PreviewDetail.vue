<template>
    <div>
        <div class="container">
            <br>
        </div>
        <div class="container">
            <router-link to="/preview">
                <button type="button" class="btn btn-success">
                    <i class="fa-solid fa-backward"></i>
                </button>
            </router-link>
        </div>
        <div class="container">
            <br>
        </div>
        <div class="container">
            <div class="row">
                <div v-if="loading" class="text-center mt-5">
                <p>Loading...</p>
                </div>
            </div>
        </div>
        <div class="container">
            <div class="row">
                <div class="col-md-12">
                    <div v-if="!loading && blog">
                        <h1 class="mt-4 text-center">{{ blog.title }}</h1>
                        <p class="text-muted text-center">Category: {{ blog.category }}</p>
                        <hr />
                        <div class="content">{{ blog.content }}</div>
                    </div>
                    <div v-if="!loading && !blog" class="text-center mt-5">
                        <p>Blog not found.</p>
                    </div>
                </div>
            </div>
        </div>
    </div>
  </template>
  
  <script>
  import axios from 'axios';
  import Swal from 'sweetalert2';
  
  export default {
    props: {
      id: {
        type: String,
        required: true,
      },
    },
    data() {
      return {
        blog: null,
        loading: true,
      };
    },
    async created() {
      try {
        const response = await axios.get(import.meta.env.VITE_API_URL + `/article/${this.id}`);
        this.blog = response.data.data;
      } catch (error) {
        if (error.response.status == 404) {
            Swal.fire({
                icon: 'Error',
                title: 'Error!',
                text: `Article not Found`,
            }).then(() => {
                this.$router.go(-1);
            });
        }
      } finally {
        this.loading = false;
      }
    },
  };
  </script>
  
<style>
.content {
  overflow-wrap: break-word; 
  word-wrap: break-word; 
}
</style>