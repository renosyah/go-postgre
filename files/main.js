new Vue({
    el: '#app',
    data() {
        return {
            users : [],
            user : {
                name : "",
                email : ""
            },
            user_edit : {
                id : "",
                name : "",
                email : ""
            },
            query : {
                search_by: "name",
                search_value: "",
                order_by: "name",
                order_dir: "asc",
                offset: 0,
                limit: 5,
                page : 1,
            },
            is_online : true,
            is_loading : false,
            host : {
                name : "",
                protocol : "",
                port : ""
            }
        }
    },
    created(){
        window.addEventListener('offline', () => { this.is_online = false })
        window.addEventListener('online', () => { this.is_online = true })
        window.history.pushState({ noBackExitsApp: true }, '')
        window.addEventListener('popstate', this.backPress)
        this.setCurrentHost()
    },
    mounted () {
        window.$('.dropdown-trigger').dropdown()
        window.$('.modal').modal({opacity:0.05,dismissible: false,preventScrolling:false})
        window.$('.sidenav').sidenav()
        this.listUser()
    },
    computed : {
        getPageName(){
            let param = new URLSearchParams(window.location.search)
            let name = param.get('page')
            return name ? name : "main-page"
        }
    },
    methods : {
        addUser(){
            this.is_loading = true
            axios({
                method: 'post',
                url:  this.baseUrl() + "api/v1/users",
                data: this.user
            }).then(response => {
                    this.is_loading = false
                    if (response.data.status == 404) {
                        window.location = this.baseUrl()
                        return
                    }
                    this.listUser()
                    this.reset()
                })
                .catch(e => {
                    console.log(e)
                    this.is_loading = false
                })
        },
        listUser(){
            this.is_loading = true
            axios({
                method: 'post',
                url:  this.baseUrl() + "api/v1/users-list",
                data: this.query
              }).then(response => {
                    this.is_loading = false
                    if (response.data.status == 404) {
                        window.location = this.baseUrl()
                        return
                    }
                    this.users = response.data.data ? response.data.data : []
                })
                .catch(e => {
                    console.log(e)
                    this.is_loading = false
                })
        },
        detailUser(id){
            this.is_loading = true
            axios.get(this.baseUrl() + "api/v1/users/" + id)
                .then(response => {
                    this.is_loading = false
                    if (response.data.status == 404) {
                        window.location = this.baseUrl()
                        return
                    }
                })
                .catch(e => {
                    console.log(e)
                    this.is_loading = false
                })
        },
        updateUser(id,data){
            this.is_loading = true
            axios({
                method: 'put',
                url:  this.baseUrl() + "api/v1/users/" + id,
                data: data
              }).then(response => {
                    this.is_loading = false
                    if (response.data.status == 404) {
                        window.location = this.baseUrl()
                        return
                    }
                    if (!response.data.data){
                        return
                    }
                    this.listUser()
                })
                .catch(e => {
                    console.log(e)
                    this.is_loading = false
                })
        },
        deleteUser(id){
            this.is_loading = true
            axios({
                method: 'delete',
                url:  this.baseUrl() + "api/v1/users/" + id,
                data: {}
              }).then(response => {
                    this.is_loading = false
                    if (response.data.status == 404) {
                        window.location = this.baseUrl()
                        return
                    }
                    this.listUser()
                })
                .catch(e => {
                    console.log(e)
                    this.is_loading = false
                })
        },
        reset(){
            this.user = {
                name : "",
                email : ""
            }
        },
        switchPage(name){
            if ('URLSearchParams' in window) {
                var searchParams = new URLSearchParams(window.location.search);
                searchParams.set('page', name);
                window.location.search = searchParams.toString();
            }
        },
        newPage(name){
            if ('URLSearchParams' in window) {
                var searchParams = new URLSearchParams(window.location.search);
                searchParams.set('page', name);
                window.open("?" + searchParams.toString(), '_blank')
            }
        },
        switchLang(lang){
            if ('URLSearchParams' in window) {
                var searchParams = new URLSearchParams(window.location.search);
                searchParams.set('lang', lang);
                window.location.search = searchParams.toString();
            }
        },
        langCheck(lang){
            let def = "ind"
            let param = new URLSearchParams(window.location.search)
            let name = param.get('lang')
            return name ? (name == lang) : (def == lang)
        },
        backPress(){
            if (event.state && event.state.noBackExitsApp) {
                window.history.pushState({ noBackExitsApp: true }, '')
            }
        },
        setCurrentHost(){
            this.host.name = window.location.hostname
            this.host.port = location.port
            this.host.protocol = location.protocol.concat("//")
        },
        baseUrl(){
            return this.host.protocol.concat(this.host.name + ":" + this.host.port + "/")
        }
    }
})

