import { ref, onMounted } from "vue"
import api from "@/api/axios"

export function useTrustedCompany() {
    const trustedCompany = ref({})
    const loading = ref(true)
    const error = ref(null)

    onMounted(async () => {
        try {
            const response = await api.get('/trusted-company')
            trustedCompany.value = response.data
        } catch (err) {
            error.value = err.message
        } finally {
            loading.value = false
        }
    })

    return { trustedCompany, loading, error }
}
