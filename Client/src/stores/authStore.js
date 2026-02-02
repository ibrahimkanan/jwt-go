import {create} from "zustand";
import axios from "axios";

const authStore = create((set) => ({
    loggedIn: null,
    user: null,
    
    login: async (loginData) => {
        try{
            await axios.post('http://localhost:3000/login', loginData, {
                withCredentials: true,
            })
            set({loggedIn: true})
        }

        catch (error){
            console.log(error)
            throw error;
        }
    },
    logout: async () => {
        try{
            await axios.post('http://localhost:3000/logout', {}, {
                withCredentials: true,
            })
            set({loggedIn: false})
        }

        catch (error){
            console.log(error)
            throw error;
        }
    },
    signup: async (signupData) => {
        try{
            await axios.post('http://localhost:3000/signup', signupData, {
                withCredentials: true,
            })
            set({loggedIn: true})
        }

        catch (error){
            console.log(error)
            throw error;
        }
    },
    checkAuth: async () => {
        try{
            await axios.get('http://localhost:3000/validate', {
                withCredentials: true,
            })
            set({loggedIn: true})
        }

        catch (error){
            set({loggedIn: false, user: null})
            console.log(error)
            throw error;
        }
    }
}))

export default authStore;