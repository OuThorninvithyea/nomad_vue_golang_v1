import { ref, onMounted } from "vue"
import api from "@/api/axios"

export function useTraveling() {
    const traveling = ref({})
    const loading = ref(true)
    const error = ref(null)

    onMounted(async () => {
        try {
            const response = await api.get('/traveling')
            traveling.value = response.data
        } catch (err) {
            error.value = err.message
        } finally {
            loading.value = false
        }
    })

    return { traveling, loading, error }
}
