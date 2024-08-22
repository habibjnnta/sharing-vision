<template>
    <div>
      <div class="container">
        <br>
      </div>
      <div class="container">
        <div class="row">
           <div class="col-md-6 d-flex justify-content-start">
            <router-link to="/preview">
                <button type="button" class="btn btn-success">
                    <i class="fa-solid fa-magnifying-glass"></i>
                </button>
            </router-link>
          </div>
          <div class="col-md-6 d-flex justify-content-end">
            <button type="button" class="btn btn-success" @click="openAddModal">
              <i class="fa-solid fa-plus"></i>
            </button>
          </div>
        </div>
      </div>
      <div class="container">
        <br>
      </div>
      <div class="container">
        <div class="row">
          <ul class="nav nav-tabs">
            <li class="nav-item" v-for="(tab, index) in tabs" :key="index">
              <a
                class="nav-link"
                :class="{ active: selectedTab === tab }"
                href="javascript:void(0);"
                @click="selectTab(tab)"
              >
                {{ tab }}
              </a>
            </li>
          </ul>
          <div class="tab-content mt-3">
            <div v-show="selectedTab === 'Publish'">
              <Publish @edit-item="openEditModal" @delete-item="handleDelete" />
            </div>
            <div v-show="selectedTab === 'Draft'">
              <Draft @edit-item="openEditModal" @delete-item="handleDelete" />
            </div>
            <div v-show="selectedTab === 'Thrash'">
              <Thrash @edit-item="openEditModal" @delete-item="handleDelete" />
            </div>
          </div>
        </div>
      </div>
  
      <!-- Modal Add -->
      <div v-if="showAddModal" class="modal fade show" tabindex="-1" style="display: block;" aria-modal="true" role="dialog">
        <div class="modal-dialog" role="document">
          <div class="modal-content">
            <div class="modal-header">
              <h5 class="modal-title">Add Article</h5>
              <button type="button" class="btn-close" @click="closeAddModal" aria-label="Close"></button>
            </div>
            <div class="modal-body">
              <form @submit.prevent="handleSubmit">
                <div class="mb-3">
                  <label for="title" class="form-label">Title</label>
                  <input type="text" class="form-control" id="title" v-model="formData.title" />
                </div>
                <div class="mb-3">
                  <label for="category" class="form-label">Category</label>
                  <input type="text" class="form-control" id="category" v-model="formData.category" />
                </div>
                <div class="mb-3">
                  <label for="content" class="form-label">Content</label>
                  <textarea class="form-control" id="content" v-model="formData.content"></textarea>
                </div>
                <div class="d-flex justify-content-end">
                  <button type="button" class="btn btn-secondary me-2" @click="handleSubmit('draft')">Draft</button>
                  <button type="submit" class="btn btn-primary" @click="handleSubmit('publish')">Publish</button>
                </div>
              </form>
            </div>
          </div>
        </div>
      </div>
  
      <!-- Modal Edit -->
      <div v-if="showEditModal" class="modal fade show" tabindex="-1" style="display: block;" aria-modal="true" role="dialog">
        <div class="modal-dialog" role="document">
          <div class="modal-content">
            <div class="modal-header">
              <h5 class="modal-title">Edit Article</h5>
              <button type="button" class="btn-close" @click="closeEditModal" aria-label="Close"></button>
            </div>
            <div class="modal-body">
              <form @submit.prevent="handleUpdate">
                <div class="mb-3">
                  <label for="editTitle" class="form-label">Title</label>
                  <input type="text" class="form-control" id="editTitle" v-model="editFormData.title" />
                </div>
                <div class="mb-3">
                  <label for="editCategory" class="form-label">Category</label>
                  <input type="text" class="form-control" id="editCategory" v-model="editFormData.category" />
                </div>
                <div class="mb-3">
                  <label for="editContent" class="form-label">Content</label>
                  <textarea class="form-control" id="editContent" v-model="editFormData.content"></textarea>
                </div>
                <div class="d-flex justify-content-end">
                    <button v-if="editFormData.status == 'publish'" type="button" class="btn btn-primary" @click="handleUpdate('draft')">Draft</button>
                    <button v-if="editFormData.status == 'draft'" type="submit" class="btn btn-primary" @click="handleUpdate('publish')">Publish</button>
                </div>
              </form>
            </div>
          </div>
        </div>
      </div>
  
      <!-- Background overlay -->
      <div v-if="showAddModal || showEditModal" class="modal-backdrop fade show"></div>
  
    </div>
  </template>
  
  <script>
  import Publish from './Tab/Publish.vue';
  import Draft from './Tab/Draft.vue';
  import Thrash from './Tab/Thrash.vue';
  import axios from 'axios';
  import Swal from 'sweetalert2';
  
  export default {
    components: {
      Publish,
      Draft,
      Thrash,
    },
  
    data() {
      return {
        tabs: ["Publish", "Draft", "Thrash"],
        selectedTab: "Publish",
        showAddModal: false,
        showEditModal: false,
        formData: { 
          title: "",
          category: "",
          content: "",
          status: "", 
        },
        editFormData: {
          id: "",
          title: "",
          category: "",
          content: "",
          status: "", 
        },
      };
    },
    methods: {
      openAddModal() {
        this.showAddModal = true;
      },
      closeAddModal() {
        this.showAddModal = false;
        this.resetForm();
      },
      openEditModal(item) {
        this.editFormData = { ...item };
        this.showEditModal = true;
      },
      closeEditModal() {
        this.showEditModal = false;
        this.resetEditForm();
      },
      async handleSubmit(action) {
        try {
          this.formData.status = action;
          await axios.post(import.meta.env.VITE_API_URL + '/article', this.formData);
          this.closeAddModal();
          Swal.fire({
            icon: 'success',
            title: 'Success!',
            text: `Article has been created successfully.`,
          }).then(() => {
            window.location.reload();
          });
        } catch (error) {
            Swal.fire({
                icon: 'error',
                title: 'Error!',
                text: error.response.statusText,
            }).then(() => {
                window.location.reload();
            });
        }
      },
      async handleUpdate(action) {
        try {
          this.editFormData.status = action;
          await axios.put(import.meta.env.VITE_API_URL + `/article/${this.editFormData.id}`, this.editFormData);
          this.closeEditModal();
          Swal.fire({
            icon: 'success',
            title: 'Success!',
            text: `Article has been updated successfully.`,
          }).then(() => {
            window.location.reload();
          });
        } catch (error) {
            Swal.fire({
                icon: 'error',
                title: 'Error!',
                text: error.response.statusText,
            }).then(() => {
                window.location.reload();
            });
        }
      },
      async handleDelete(id) {
        Swal.fire({
            title: 'Are you sure?',
            text: "You won't be able to revert this!",
            icon: 'warning',
            showCancelButton: true,
            confirmButtonColor: '#3085d6',
            cancelButtonColor: '#d33',
            confirmButtonText: 'Yes, delete it!'
        }).then(async (result) => {
            if (result.isConfirmed) {
                try {
                    await axios.delete(import.meta.env.VITE_API_URL + `/article/${id}`);
                    Swal.fire(
                        'Deleted!',
                        'Your article has been deleted.',
                        'success'
                    ).then(() => {
                        window.location.reload();
                    });
                } catch (error) {
                    Swal.fire(
                        'Error!',
                        'Failed to delete the article.',
                        'error'
                    ).then(() => {
                        window.location.reload();
                    });
                }
            }
        });
      },
      resetForm() {
        this.formData = { 
          title: "",
          category: "",
          content: "",
          status: "", 
        };
      },
      resetEditForm() {
        this.editFormData = {
          id: "",
          title: "",
          category: "",
          content: "",
        };
      },
      selectTab(tab) {
        this.selectedTab = tab;
      }
    },
  };
  </script>
  
  <style scoped>
  /* Add any component-specific styles here */
  </style>
  