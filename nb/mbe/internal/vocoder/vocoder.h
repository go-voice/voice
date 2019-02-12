#ifndef _IMBE_C_H
#define _IMBE_C_H

#include <stdint.h>

#ifdef __cplusplus
#include "vocoder_mbe.h"
extern "C"
{
#endif

#ifdef __cplusplus
typedef mbe_vocoder *vocoder;
#else
typedef void *vocoder;
#endif

vocoder vocoder_new(int quality);
void vocoder_destroy(vocoder);
void vocoder_reset(vocoder);
void vocoder_ambe_mode_dmr(vocoder coder);
void vocoder_ambe_mode_dstar(vocoder coder);
void vocoder_ambe_encode(vocoder coder, uint8_t *dst, int16_t *src);
void vocoder_ambe_decode(vocoder coder, int16_t *dst, uint8_t *src);
void vocoder_imbe_encode(vocoder coder, int16_t *frame_vector, int16_t *snd);
void vocoder_imbe_decode(vocoder coder, int16_t *frame_vector, int16_t *snd);

#ifdef __cplusplus
}
#endif

#endif