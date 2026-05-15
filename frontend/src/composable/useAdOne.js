import { ref, onMounted } from "vue"
import api from "@/api/axios"

export function useAdOne() {
    const adOne = ref({})
    const loading = ref(true)
    const error = ref(null)

    onMounted(async () => {
        try {
            const response = await api.get('/ad-one')
            adOne.value = response.data
        } catch (err) {
            error.value = err.message
        } finally {
            loading.value = false
        }
    })

    return { adOne, loading, error }
}
