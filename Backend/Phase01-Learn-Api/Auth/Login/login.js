// Dummy user data
var users = [
    { username: "admin", password: "admin123" },
    { username: "user", password: "user123" }
];
// Login function
function login(username, password) {
    var user = users.find(function (u) { return u.username === username && u.password === password; });
    return !!user;
}
// Contoh penggunaan
var usernameInput = "admin";
var passwordInput = "admin123";
if (login(usernameInput, passwordInput)) {
    console.log("Login berhasil!");
}
else {
    console.log("Login gagal!");
}
