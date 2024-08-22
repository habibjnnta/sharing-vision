<template>
    <div>
        <div class="container">
            <br>
        </div>

        <div class="container">
            <router-link to="/">
                <button type="button" class="btn btn-success">
                    <i class="fa-solid fa-backward"></i>
                </button>
            </router-link>
        </div>

        <div class="container">
            <br>
        </div>

        <div class="container">
            <div class="row d-flex justify-content-center text-center">
                <div class="col-md-3" v-for="item in items">
                    <div class="card" style="width: 18rem;">
                        <div class="card-body">
                            <h5 class="card-title">{{ item.title }}</h5>
                            <p class="card-text">
                                {{ item.content.length > 50 ? item.content.substring(0, 50) + "..." : item.content }}
                            </p>
                            <router-link :to="{ name: 'Preview Article Detail', params: { id: item.id } }">
                                <a href="javascript:void(0);" class="card-link">Read more...</a>
                            </router-link>
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <div class="container">
            <br>
        </div>

        <div class="container" v-if="items.length > 0">
            <div class="row">
                <div class="col-md-12 d-flex justify-content-center">
                    <!-- Pagination Controls -->
                    <nav aria-label="Page navigation">
                        <ul class="pagination">
                            <li class="page-item" :class="{ disabled: meta.page == meta.previous_page }">
                                <button class="page-link" @click="goToPage(meta.previous_page)" :disabled="meta.page == meta.previous_page">Previous</button>
                            </li>
                            <li
                                v-for="page in meta.total_page"
                                :key="page"
                                class="page-item"
                                :class="{ active: page === currentPage }"
                            >
                                <button class="page-link" @click="goToPage(page)">{{ page }}</button>
                            </li>
                            <li class="page-item" :class="{ disabled: meta.next_page > meta.total_page || meta.count < meta.limit }">
                                <button class="page-link" @click="goToPage(meta.next_page)" :disabled="meta.next_page > meta.total_page || meta.count < meta.limit">Next</button>
                            </li>
                        </ul>
                    </nav>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    import axios from 'axios';

    export default {
        data() {
            return {
                items: [],
                meta: {
                    limit: import.meta.env.VITE_APP_LIMIT,
                    page: 1,
                    next_page: 0,
                    previous_page: 0,
                    count: 0,
                    total_page: 0
                },
                currentPage: 1,
                currentLimit: import.meta.env.VITE_APP_LIMIT,
            }
        },

        mounted() {
            this.getData(this.currentPage, this.currentLimit);
        },

        methods: {
            async getData(page, limit) {
                try {
                    const response = await axios.get(`${import.meta.env.VITE_API_URL}/article?status=publish&limit=${limit}&page=${page}`);
                    this.items = response.data.data;
                    let data = response.data.meta
                    if (data.page != 0) {
                        this.meta.page = data.page
                    }
                    if (data.next_page != 0) {
                        this.meta.next_page = data.next_page
                    }
                    if (data.previous_page != 0) {
                        this.meta.previous_page = data.previous_page
                    }
                    if (data.count != 0) {
                        this.meta.count = data.count
                    }
                    if (data.total_page != 0) {
                        this.meta.total_page = data.total_page
                    }
                    console.log(this.meta);
                    this.currentPage = page;
                    this.currentLimit = limit;
                } catch (error) {
                    console.error('Error fetching data:', error);
                }
            },
        
            goToPage(page) {
                if (page > 0 && page <= this.meta.total_page) {
                    this.getData(page, this.meta.limit);
                }
            },
        }
    }
</script>