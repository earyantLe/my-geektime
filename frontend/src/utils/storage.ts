export const storage = {
  getToken: () => localStorage.getItem('token'),
  setToken: (token: string) => localStorage.setItem('token', token),
  removeToken: () => localStorage.removeItem('token'),
  
  getUser: () => {
    const uid = localStorage.getItem('uid')
    const uname = localStorage.getItem('uname')
    const avatar = localStorage.getItem('avatar')
    const roleId = localStorage.getItem('role_id')
    return { uid, uname, avatar, roleId }
  },
  
  setUser: (user: { uid: string; user_name: string; avatar: string; role_id: number }) => {
    localStorage.setItem('uid', user.uid)
    localStorage.setItem('uname', user.user_name)
    localStorage.setItem('avatar', user.avatar)
    localStorage.setItem('role_id', String(user.role_id))
  },
  
  removeUser: () => {
    localStorage.removeItem('uid')
    localStorage.removeItem('uname')
    localStorage.removeItem('avatar')
    localStorage.removeItem('role_id')
  },
  
  getGeekAuth: () => localStorage.getItem('geek_auth'),
  setGeekAuth: (value: string) => localStorage.setItem('geek_auth', value),
  removeGeekAuth: () => localStorage.removeItem('geek_auth'),
  
  clear: () => localStorage.clear(),
}
