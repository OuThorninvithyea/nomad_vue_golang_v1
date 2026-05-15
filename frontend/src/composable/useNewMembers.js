import { ref, onMounted } from "vue"
import api from "@/api/axios"

export function useNewMembers() {
    const newMembers = ref({})
    const loading = ref(true)
    const error = ref(null)

    onMounted(async () => {
        try {
            const response = await api.get('/new-members')
            newMembers.value = response.data
        } catch (err) {
            error.value = err.message
        } finally {
            loading.value = false
        }
    })

    return { newMembers, loading, error }
}
