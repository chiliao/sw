#include "nic/hal/src/qos.hpp"
#include "nic/include/pd.hpp"
#include "nic/include/qos_api.hpp"

namespace hal {

// ----------------------------------------------------------------------------
// Qos-class API: Get Qos-class's hal handle
// ----------------------------------------------------------------------------
hal_handle_t 
qos_class_get_qos_class_handle(qos_class_t *pi_qos_class)
{
    return pi_qos_class->hal_handle;
}

// ----------------------------------------------------------------------------
// Qos-class API: Set PD Qos-class in PI Qos-class
// ----------------------------------------------------------------------------
void 
qos_class_set_pd_qos_class(qos_class_t *pi_qos_class, pd::pd_qos_class_t *pd_qos_class)
{
    pi_qos_class->pd = pd_qos_class;
}

// ----------------------------------------------------------------------------
// Qos-class API: Get PD Qos-class in PI Qos-class
// ----------------------------------------------------------------------------
pd::pd_qos_class_t* 
qos_class_get_pd_qos_class(qos_class_t *pi_qos_class)
{
    return pi_qos_class->pd;
}

qos_group_t 
qos_class_get_qos_group (qos_class_t *pi_qos_class)
{
    return pi_qos_class->key.qos_group;
}

bool 
qos_class_is_no_drop (qos_class_t *pi_qos_class)
{
    return pi_qos_class->no_drop;
}

} // namespace hal
