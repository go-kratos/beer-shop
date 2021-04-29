import service from './index'

function login(username, password) {
    return service.post("/login", {
        username: username,
        password: password,
    })
}

function register(username, password) {
    return service.post("/register", {
        username: username,
        password: password,
    })
}

function logout(username, password) {
    return service.post("/logout", {
        username: username,
        password: password,
    })
}
