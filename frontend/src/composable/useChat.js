import { ref, onMounted } from "vue"
import api from "@/api/axios"

export function useChat() {
    const chat = ref({})
    const loading = ref(true)
    const error = ref(null)

    onMounted(async () => {
        try {
            const response = await api.get('/chat')
            chat.value = response.data
        } catch (err) {
            error.value = err.message
        } finally {
            loading.value = false
        }
    })

    return { chat, loading, error }
}
