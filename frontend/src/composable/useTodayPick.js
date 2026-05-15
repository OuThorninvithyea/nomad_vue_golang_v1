import { ref, onMounted } from "vue"
import api from "@/api/axios"

export function useTodayPick() {
    const todayPick = ref({})
    const loading = ref(true)
    const error = ref(null)

    onMounted(async () => {
        try {
            const response = await api.get('/today-pick')
            todayPick.value = response.data
        } catch (err) {
            error.value = err.message
        } finally {
            loading.value = false
        }
    })

    return { todayPick, loading, error }
}
