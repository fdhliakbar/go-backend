interface User {
  username: string;
  password: string;
}

const users: User[] = [
  { username: "admin", password: "admin123" },
  { username: "user", password: "user123" }
];

function login(username: string, password: string): boolean {
  const user = users.find(u => u.username === username && u.password === password);
  return !!user;
}


const usernameInput = "admin";
const passwordInput = "admin123";

if (login(usernameInput, passwordInput)) {
  console.log("Login berhasil!");
} else {
  console.log("Login gagal!");
}

console.log("Pengguna yang tersedia:");