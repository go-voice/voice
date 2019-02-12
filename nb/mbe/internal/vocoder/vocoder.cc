#include "vocoder.h"
#include "mbelib.h"

mbe_vocoder::mbe_vocoder(int uvquality)
    : dmr_mode(false),
      dstar_mode(false),
      uvquality(uvquality)
{
    ambe = new ambe_encoder();
	ambe->set_gain_adjust(-1);
    imbe = new imbe_vocoder();
    reset();
}

mbe_vocoder::~mbe_vocoder(void)
{
    delete(ambe);
    delete(imbe);
}

static void pack_49bit(uint8_t dst[8], const char bits[49])
{
	dst[0]  = (bits[ 0] & 1) << 7;
	dst[0] |= (bits[ 1] & 1) << 6;
	dst[0] |= (bits[ 2] & 1) << 5;
	dst[0] |= (bits[ 3] & 1) << 4;
	dst[0] |= (bits[ 4] & 1) << 3;
	dst[0] |= (bits[ 5] & 1) << 2;
	dst[0] |= (bits[ 6] & 1) << 1;
	dst[0] |= (bits[ 7] & 1);
	dst[1]  = (bits[ 8] & 1) << 7;
	dst[1] |= (bits[ 9] & 1) << 6;
	dst[1] |= (bits[10] & 1) << 5;
	dst[1] |= (bits[11] & 1) << 4;
	dst[1] |= (bits[12] & 1) << 3;
	dst[1] |= (bits[13] & 1) << 2;
	dst[1] |= (bits[14] & 1) << 1;
	dst[1] |= (bits[15] & 1);
	dst[2]  = (bits[16] & 1) << 7;
	dst[2] |= (bits[17] & 1) << 6;
	dst[2] |= (bits[18] & 1) << 5;
	dst[2] |= (bits[19] & 1) << 4;
	dst[2] |= (bits[20] & 1) << 3;
	dst[2] |= (bits[21] & 1) << 2;
	dst[2] |= (bits[22] & 1) << 1;
	dst[2] |= (bits[23] & 1);
	dst[3]  = (bits[24] & 1) << 7;
	dst[3] |= (bits[25] & 1) << 6;
	dst[3] |= (bits[26] & 1) << 5;
	dst[3] |= (bits[27] & 1) << 4;
	dst[3] |= (bits[28] & 1) << 3;
	dst[3] |= (bits[29] & 1) << 2;
	dst[3] |= (bits[30] & 1) << 1;
	dst[3] |= (bits[31] & 1);
	dst[4]  = (bits[32] & 1) << 7;
	dst[4] |= (bits[33] & 1) << 6;
	dst[4] |= (bits[34] & 1) << 5;
	dst[4] |= (bits[35] & 1) << 4;
	dst[4] |= (bits[36] & 1) << 3;
	dst[4] |= (bits[37] & 1) << 2;
	dst[4] |= (bits[38] & 1) << 1;
	dst[4] |= (bits[30] & 1);
	dst[5]  = (bits[40] & 1) << 7;
	dst[5] |= (bits[41] & 1) << 6;
	dst[5] |= (bits[42] & 1) << 5;
	dst[5] |= (bits[43] & 1) << 4;
	dst[5] |= (bits[44] & 1) << 3;
	dst[5] |= (bits[45] & 1) << 2;
	dst[5] |= (bits[46] & 1) << 1;
	dst[5] |= (bits[47] & 1);
	dst[6]  = (bits[48] & 1) << 7;
}

void mbe_vocoder::ambe_encode(uint8_t *dst, int16_t *src)
{
	if (dmr_mode)
	{
		char ambe_d[49];
		ambe->encode(src, (uint8_t *)ambe_d);
		pack_49bit(dst, ambe_d);
	}
}

static void decode_49bit(char bits[49], const uint8_t src[8])
{
    bits[ 0] = (src[0] & 0x80) >> 7;
	bits[ 1] = (src[0] & 0x40) >> 6;
	bits[ 2] = (src[0] & 0x20) >> 5;
	bits[ 3] = (src[0] & 0x10) >> 4;
	bits[ 4] = (src[0] & 0x08) >> 3;
	bits[ 5] = (src[0] & 0x04) >> 2;
	bits[ 6] = (src[0] & 0x02) >> 1;
	bits[ 7] = (src[0] & 0x01);
	bits[ 8] = (src[1] & 0x80) >> 7;
	bits[ 9] = (src[1] & 0x40) >> 6;
	bits[10] = (src[1] & 0x20) >> 5;
	bits[11] = (src[1] & 0x10) >> 4;
	bits[12] = (src[1] & 0x08) >> 3;
	bits[13] = (src[1] & 0x04) >> 2;
	bits[14] = (src[1] & 0x02) >> 1;
	bits[15] = (src[1] & 0x01);
	bits[16] = (src[2] & 0x80) >> 7;
	bits[17] = (src[2] & 0x40) >> 6;
	bits[18] = (src[2] & 0x20) >> 5;
	bits[19] = (src[2] & 0x10) >> 4;
	bits[20] = (src[2] & 0x08) >> 3;
	bits[21] = (src[2] & 0x04) >> 2;
	bits[22] = (src[2] & 0x02) >> 1;
	bits[23] = (src[2] & 0x01);
	bits[24] = (src[3] & 0x80) >> 7;
	bits[25] = (src[3] & 0x40) >> 6;
	bits[26] = (src[3] & 0x20) >> 5;
	bits[27] = (src[3] & 0x10) >> 4;
	bits[28] = (src[3] & 0x08) >> 3;
	bits[29] = (src[3] & 0x04) >> 2;
	bits[30] = (src[3] & 0x02) >> 1;
	bits[31] = (src[3] & 0x01);
	bits[32] = (src[4] & 0x80) >> 7;
	bits[33] = (src[4] & 0x40) >> 6;
	bits[34] = (src[4] & 0x20) >> 5;
	bits[35] = (src[4] & 0x10) >> 4;
	bits[36] = (src[4] & 0x08) >> 3;
	bits[37] = (src[4] & 0x04) >> 2;
	bits[38] = (src[4] & 0x02) >> 1;
	bits[39] = (src[4] & 0x01);
	bits[40] = (src[5] & 0x80) >> 7;
	bits[41] = (src[5] & 0x40) >> 6;
	bits[42] = (src[5] & 0x20) >> 5;
	bits[43] = (src[5] & 0x10) >> 4;
	bits[44] = (src[5] & 0x08) >> 3;
	bits[45] = (src[5] & 0x04) >> 2;
	bits[46] = (src[5] & 0x02) >> 1;
	bits[47] = (src[5] & 0x01);
	bits[48] = (src[6] & 0x80) >> 7;
}

void mbe_vocoder::ambe_decode(int16_t *dst, uint8_t *src)
{
    if (dmr_mode)
    {
        char ambe_d[49];
        decode_49bit(ambe_d, src);
        mbe_processAmbe2450Data(
			dst, &this->errs, &this->errs2, this->err_str,
			ambe_d, &this->cur_mp, &this->prev_mp, &this->enh_mp,
			this->uvquality);
    }
}

void mbe_vocoder::ambe_set_49bit(void)
{
    ambe->set_49bit_mode();
    dmr_mode = true;
}

void mbe_vocoder::ambe_set_dstar(void)
{
    ambe->set_dstar_mode();
    dstar_mode = true;
}

void mbe_vocoder::ambe_set_gain(float gain_adjust)
{
    ambe->set_gain_adjust(gain_adjust);
}

void mbe_vocoder::imbe_encode(int16_t *frame, int16_t *samples)
{
    imbe->imbe_encode(frame, samples);
}

void mbe_vocoder::imbe_decode(int16_t *frame, int16_t *samples)
{
    imbe->imbe_decode(frame, samples);
}

void mbe_vocoder::reset(void)
{
    mbe_initMbeParms(&cur_mp, &prev_mp, &enh_mp);
    errs = 0;
    errs2 = 0;
    err_str[0] = 0;
}

vocoder vocoder_new(int quality)
{
    return new mbe_vocoder(quality);
}

void vocoder_destroy(vocoder coder)
{
    delete(coder);
}

void vocoder_reset(vocoder coder)
{
    coder->reset();
}

void vocoder_ambe_mode_dmr(vocoder coder)
{
    coder->ambe_set_49bit();
}

void vocoder_ambe_mode_dstar(vocoder coder)
{
    coder->ambe_set_dstar();
}

void vocoder_ambe_encode(vocoder coder, uint8_t *dst, int16_t *src)
{
    coder->ambe_encode(dst, src);
}

void vocoder_ambe_decode(vocoder coder, int16_t *dst, uint8_t *src)
{
    coder->ambe_decode(dst, src);
}

void vocoder_imbe_encode(vocoder coder, int16_t *frame_vector, int16_t *snd)
{
    coder->imbe_encode(frame_vector, snd);
}

void vocoder_imbe_decode(vocoder coder, int16_t *frame_vector, int16_t *snd)
{
    coder->imbe_decode(frame_vector, snd);
}

void vocoder_gain_adjust(vocoder coder, float gain)
{
    coder->ambe_set_gain(gain);
}
