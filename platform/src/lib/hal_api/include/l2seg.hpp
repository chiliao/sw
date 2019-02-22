
#ifndef __L2SEG_HPP__
#define __L2SEG_HPP__

#include <vector>

#include "lib/indexer/indexer.hpp"

#include "hal_types.hpp"
#include "hal.hpp"
#include "vrf.hpp"


// (vrf, vlan)
typedef std::tuple<uint32_t, vlan_t> l2seg_key_t;

// uplink_id -> Uplink
typedef std::map<uplink_id_t, Uplink*> uplink_map_t;


class HalL2Segment : public HalObject
{
public:
    static HalL2Segment *Factory(HalVrf *vrf, uint16_t vlan);
    static hal_irisc_ret_t Destroy(HalL2Segment *l2seg);

    hal_irisc_ret_t HalL2SegmentCreate();
    hal_irisc_ret_t HalL2SegmentDelete();

    static HalL2Segment *Lookup(HalVrf *vrf, uint16_t vlan);


    hal_irisc_ret_t AddUplink(Uplink *uplink);
    hal_irisc_ret_t DelUplink(Uplink *uplink);

    uint64_t GetId();
    HalVrf *GetVrf();

    static void Probe();

private:
  HalL2Segment(HalVrf *vrf, uint16_t vlan_id);
  ~HalL2Segment() {};
  hal_irisc_ret_t TriggerHalUpdate();

  uint32_t id;
  vlan_t _vlan;
  HalVrf *vrf;
  uplink_map_t uplink_refs;

  // L2seg id
  static sdk::lib::indexer *allocator;
  static constexpr uint64_t max_l2segs = 4096;

  static std::map<l2seg_key_t, HalL2Segment*> l2seg_db;

};

#endif /* __L2SEG_HPP__ */
