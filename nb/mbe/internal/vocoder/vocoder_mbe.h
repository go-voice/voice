#ifndef _VOCODER_MBE_H
#define _VOCODER_MBE_H

#include <stdint.h>
#include "imbe_vocoder.h"
#include "ambe_encoder.h"

class mbe_vocoder
{
  public:
    mbe_vocoder(int uvquality);
    ~mbe_vocoder(void);

    void ambe_encode(uint8_t *dst, int16_t *src);
    void ambe_decode(int16_t *dst, uint8_t *src);
    void ambe_set_49bit(void);
    void ambe_set_dstar(void);
    void ambe_set_gain(float gain_adjust);
    void imbe_encode(int16_t *frame, int16_t *samples);
    void imbe_decode(int16_t *frame, int16_t *samples);
    void reset(void);

  private:
    imbe_vocoder *imbe;
    ambe_encoder *ambe;
    int uvquality;
    mbe_parms cur_mp;
    mbe_parms prev_mp;
    mbe_parms enh_mp;
    int errs;
    int errs2;
    char err_str[64];
    bool dmr_mode;
    bool dstar_mode;
};

#endif