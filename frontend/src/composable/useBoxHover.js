import { ref, onMounted } from "vue"
import api from "@/api/axios"

export function useBoxHover() {
    const boxHover = ref({})
    const loading = ref(true)
    const error = ref(null)

    onMounted(async () => {
        try {
            const response = await api.get('/box-hover')
            boxHover.value = response.data
        } catch (err) {
            error.value = err.message
        } finally {
            loading.value = false
        }
    })

    return { boxHover, loading, error }
}
