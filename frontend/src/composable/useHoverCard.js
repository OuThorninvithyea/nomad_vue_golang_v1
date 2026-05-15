import { ref, onMounted } from "vue"
import api from "@/api/axios"

export function useHoverCard() {
    const hoverCards = ref([])
    const loading = ref(true)
    const error = ref(null)

    onMounted(async () => {
        try {
            const response = await api.get('/hover-card')
            hoverCards.value = response.data
        } catch (err) {
            error.value = err.message
        } finally {
            loading.value = false
        }
    })

    return { hoverCards, loading, error }
}
