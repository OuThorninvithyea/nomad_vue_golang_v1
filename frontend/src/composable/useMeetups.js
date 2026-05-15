import { ref, onMounted } from "vue"
import api from "@/api/axios"

export function useMeetups() {
    const meetups = ref({})
    const loading = ref(true)
    const error = ref(null)

    onMounted(async () => {
        try {
            const response = await api.get('/meetups')
            meetups.value = response.data
        } catch (err) {
            error.value = err.message
        } finally {
            loading.value = false
        }
    })

    return { meetups, loading, error }
}
