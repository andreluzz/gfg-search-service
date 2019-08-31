<template>
    <div class="container">
        <div class="box" style="margin-top:25px;">
           <section>
               <b-field grouped>
                    <b-input v-model="queryFilter" placeholder="Search products..." expanded></b-input>
                    <b-field label="Brand" label-position="on-border">
                        <b-select v-model="brandFilter" placeholder="Select a brand">
                            <option value="">All</option>
                            <option value="brand:adidas">Adidas</option>
                            <option value="brand:nike">Nike</option>
                            <option value="brand:puma">Puma</option>
                            <option value="brand:under armor">Under Armor</option>
                        </b-select>
                    </b-field>
                    <b-field label="Items" label-position="on-border">
                        <b-select v-model="itemsPerPageFilter">
                            <option value="5">5</option>
                            <option value="10">10</option>
                            <option value="20" selected>20</option>
                        </b-select>
                    </b-field>
                    <p class="control">
                        <button class="button is-primary" @click="onSearch">
                            <b-icon icon="magnify"></b-icon>
                            <span>Search</span>
                        </button>
                    </p>
                </b-field>
            </section>
        </div>
        <section>
            <b-table
                :data="data"
                :loading="loading"
                paginated
                backend-pagination
                :total="total"
                :per-page="perPage"
                @page-change="onPageChange"
                aria-next-label="Next page"
                aria-previous-label="Previous page"
                aria-page-label="Page"
                aria-current-label="Current page"
                backend-sorting
                :default-sort-direction="defaultSortOrder"
                :default-sort="[sortField, sortOrder]"
                @sort="onSort">

                <template slot-scope="props">
                    <b-table-column field="title" label="Title" sortable>
                        {{ props.row.title }}
                    </b-table-column>

                    <b-table-column field="brand" label="Brand" sortable>
                        {{ props.row.brand }}
                    </b-table-column>

                    <b-table-column field="price" label="Price" numeric sortable>
                        {{ props.row.price }}
                    </b-table-column>

                    <b-table-column field="stock" label="Stock" numeric sortable>
                        <span class="tag" :class="type(props.row.stock)">
                            {{ props.row.stock }}
                        </span>
                    </b-table-column>
                </template>
            </b-table>
        </section>
    </div>
</template>

<script>
    export default {
        data() {
            return {
                data: [],
                total: 0,
                loading: false,
                queryFilter: '',
                brandFilter: '',
                itemsPerPageFilter: 10,
                sortField: 'title',
                sortOrder: 'asc',
                defaultSortOrder: 'desc',
                page: 1,
                perPage: 10
            }
        },
        methods: {
            loadAsyncData() {
                const params = [
                    `q=${this.queryFilter}`,
                    `filter=${this.brandFilter}`,
                    `sort=${this.sortField}:${this.sortOrder}`,
                    `limit=${this.perPage}`,                
                    `page=${this.page}`
                ].join('&')

                this.loading = true
                const token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.e30.wDWyyGem9YgXDDbH3Un7YYcTB8IcN_BE4BMmS1tvlnE"
                this.$http.get(`${process.env.VUE_APP_API_URL}/products?${params}`, {headers: {'Authorization': token, 'x-service-version': 'v2'}})
                    .then(({ data }) => {
                        this.data = []
                        data.results.forEach((item) => {
                            this.data.push(item)
                        })
                        this.total = data.total
                        this.loading = false
                    })
                    .catch((error) => {
                        this.data = []
                        this.total = 0
                        this.loading = false
                        this.errorSnackBar(error)
                        throw error
                    })
            },
            onPageChange(page) {
                this.page = page
                this.loadAsyncData()
            },
            onSort(field, order) {
                this.sortField = field
                this.sortOrder = order
                this.loadAsyncData()
            },
            onSearch() {
                this.page = 1
                this.perPage = this.itemsPerPageFilter
                this.loadAsyncData()
            },
            type(value) {
                const number = parseFloat(value)
                if (number < 5) {
                    return 'is-danger'
                } else if (number >= 5 && number < 10) {
                    return 'is-warning'
                } else if (number >= 10) {
                    return 'is-success'
                }
            },
            errorSnackBar(err) {
                let errText = err.bodyText
                if (err.bodyText == '') {
                    errText = `Unavailable Rest API Server: ${process.env.VUE_APP_API_URL}`
                }
                const params = {
                    duration: 5000,
                    message: errText,
                    type: 'is-danger',
                    position: 'is-bottom',
                    queue: false
                }
                this.$buefy.snackbar.open(params)
            }
        },
        mounted() {
            this.loadAsyncData()
        }
    }
</script>