package DayOne

import (
	"testing"
)

func main() {
	print(DayOne(getInput()))
}

func DayOne(input []uint) uint {
	if len(input) == 0 {
		return 0
	}
	if len(input) == 1 {
		return 0
	}

	firstIter := true
	var prev uint = 0
	var increases uint = 0

	for _, i := range input {
		if firstIter {
			firstIter = false
			prev = i
			continue
		}
		if i > prev {
			increases += 1
		}
		prev = i
	}

	return increases
}

func getInput() []uint {
	return []uint{155, 157, 156, 172, 170, 186, 198, 189, 207, 213, 222, 228, 229, 227, 220, 226, 241, 243, 244, 246, 256, 255, 260, 266, 268, 270, 269, 271, 272, 275, 276, 277, 272, 273, 278, 286, 293, 298, 304, 305, 317, 330, 342, 341, 371, 374, 376, 377, 381, 385, 398, 409, 433, 434, 435, 436, 437, 443, 445, 455, 450, 447, 462, 465, 466, 468, 474, 493, 494, 523, 532, 539, 538, 539, 542, 547, 545, 557, 558, 561, 581, 582, 583, 580, 584, 603, 617, 618, 605, 616, 618, 620, 636, 637, 639, 640, 641, 642, 641, 642, 644, 645, 664, 668, 678, 679, 681, 686, 687, 688, 689, 706, 709, 724, 727, 735, 737, 749, 753, 754, 760, 766, 770, 765, 766, 773, 775, 777, 785, 786, 787, 797, 798, 818, 835, 836, 838, 859, 856, 860, 867, 868, 870, 871, 892, 896, 897, 898, 919, 928, 944, 943, 965, 980, 982, 988, 994, 997, 998, 997, 998, 982, 990, 986, 985, 986, 987, 997, 1003, 998, 1005, 1006, 1009, 1008, 1015, 1021, 1023, 1024, 1021, 1041, 1042, 1044, 1052, 1053, 1056, 1053, 1056, 1057, 1058, 1060, 1053, 1054, 1059, 1056, 1057, 1068, 1081, 1093, 1094, 1099, 1103, 1101, 1103, 1125, 1127, 1128, 1134, 1130, 1133, 1137, 1141, 1142, 1144, 1164, 1166, 1180, 1178, 1177, 1182, 1183, 1180, 1182, 1197, 1198, 1203, 1204, 1206, 1207, 1208, 1217, 1241, 1250, 1249, 1250, 1252, 1253, 1257, 1285, 1286, 1290, 1302, 1330, 1332, 1333, 1334, 1336, 1334, 1331, 1332, 1344, 1345, 1354, 1358, 1357, 1360, 1362, 1376, 1378, 1379, 1381, 1379, 1381, 1380, 1379, 1383, 1393, 1403, 1407, 1397, 1403, 1408, 1413, 1415, 1419, 1427, 1430, 1438, 1441, 1446, 1476, 1502, 1495, 1497, 1498, 1499, 1500, 1502, 1503, 1505, 1526, 1552, 1570, 1572, 1577, 1579, 1574, 1575, 1585, 1605, 1624, 1626, 1631, 1646, 1647, 1649, 1681, 1682, 1684, 1687, 1690, 1691, 1694, 1718, 1719, 1723, 1731, 1721, 1723, 1724, 1742, 1743, 1748, 1753, 1761, 1762, 1765, 1772, 1774, 1773, 1781, 1784, 1789, 1794, 1795, 1803, 1807, 1813, 1829, 1828, 1830, 1844, 1847, 1850, 1854, 1855, 1850, 1848, 1853, 1878, 1909, 1910, 1912, 1917, 1937, 1939, 1943, 1941, 1952, 1959, 1963, 1971, 1974, 1989, 2007, 2009, 2005, 2006, 2010, 2013, 2016, 2027, 2044, 2030, 2031, 2032, 2035, 2036, 2047, 2050, 2052, 2061, 2071, 2077, 2080, 2082, 2094, 2098, 2119, 2111, 2112, 2127, 2139, 2140, 2141, 2149, 2151, 2156, 2157, 2171, 2186, 2202, 2206, 2221, 2239, 2247, 2252, 2262, 2264, 2271, 2281, 2291, 2293, 2333, 2335, 2339, 2343, 2338, 2342, 2337, 2343, 2344, 2347, 2353, 2358, 2363, 2364, 2377, 2403, 2402, 2405, 2418, 2408, 2406, 2422, 2424, 2422, 2435, 2437, 2439, 2446, 2458, 2459, 2484, 2486, 2487, 2484, 2486, 2485, 2508, 2509, 2507, 2508, 2509, 2521, 2531, 2541, 2542, 2545, 2556, 2566, 2579, 2582, 2583, 2584, 2585, 2591, 2593, 2594, 2598, 2599, 2600, 2601, 2604, 2609, 2616, 2638, 2647, 2646, 2645, 2642, 2643, 2644, 2645, 2624, 2635, 2634, 2639, 2643, 2645, 2658, 2659, 2662, 2666, 2654, 2655, 2666, 2663, 2673, 2684, 2702, 2707, 2708, 2734, 2742, 2745, 2748, 2751, 2754, 2753, 2751, 2758, 2757, 2775, 2761, 2772, 2773, 2796, 2807, 2808, 2809, 2810, 2817, 2819, 2823, 2824, 2847, 2850, 2851, 2859, 2861, 2862, 2854, 2860, 2877, 2880, 2881, 2880, 2913, 2916, 2917, 2918, 2938, 2954, 2957, 2965, 2971, 2976, 2984, 2985, 2986, 3002, 3003, 3004, 3005, 2990, 3004, 2997, 2995, 3011, 3005, 3009, 3018, 3019, 3021, 3015, 3016, 3017, 3024, 3026, 3041, 3042, 3052, 3064, 3063, 3068, 3087, 3088, 3084, 3110, 3109, 3121, 3131, 3130, 3134, 3139, 3145, 3153, 3159, 3162, 3179, 3219, 3224, 3220, 3221, 3230, 3232, 3256, 3236, 3235, 3237, 3245, 3250, 3265, 3222, 3238, 3273, 3287, 3288, 3293, 3297, 3307, 3292, 3306, 3311, 3312, 3324, 3315, 3316, 3318, 3316, 3317, 3318, 3320, 3321, 3328, 3330, 3331, 3333, 3337, 3342, 3323, 3332, 3330, 3337, 3338, 3347, 3343, 3345, 3347, 3353, 3361, 3362, 3363, 3366, 3351, 3353, 3365, 3368, 3389, 3400, 3410, 3420, 3438, 3439, 3445, 3444, 3472, 3486, 3484, 3494, 3503, 3504, 3505, 3513, 3524, 3526, 3527, 3535, 3536, 3529, 3530, 3531, 3533, 3547, 3549, 3562, 3553, 3560, 3561, 3562, 3566, 3567, 3568, 3578, 3579, 3581, 3611, 3613, 3615, 3618, 3639, 3655, 3657, 3676, 3680, 3700, 3709, 3710, 3717, 3720, 3727, 3735, 3736, 3746, 3768, 3772, 3777, 3780, 3790, 3799, 3802, 3804, 3812, 3814, 3815, 3823, 3824, 3827, 3853, 3861, 3854, 3846, 3849, 3868, 3879, 3883, 3880, 3881, 3889, 3891, 3890, 3899, 3891, 3885, 3917, 3918, 3919, 3920, 3926, 3921, 3931, 3938, 3939, 3950, 3951, 3955, 3966, 3965, 3970, 3982, 3996, 3997, 4000, 4002, 4003, 4005, 4010, 4015, 4023, 4027, 4019, 4025, 4034, 4035, 4037, 4053, 4075, 4084, 4089, 4093, 4083, 4089, 4090, 4129, 4141, 4158, 4185, 4189, 4191, 4201, 4202, 4212, 4222, 4223, 4238, 4244, 4243, 4245, 4252, 4270, 4275, 4278, 4279, 4283, 4285, 4289, 4290, 4293, 4294, 4302, 4306, 4290, 4292, 4310, 4317, 4318, 4321, 4322, 4325, 4303, 4304, 4306, 4307, 4310, 4330, 4328, 4329, 4330, 4332, 4335, 4337, 4336, 4344, 4347, 4342, 4361, 4362, 4361, 4359, 4379, 4385, 4390, 4389, 4392, 4395, 4397, 4412, 4415, 4416, 4418, 4441, 4442, 4443, 4441, 4435, 4437, 4450, 4471, 4472, 4473, 4469, 4470, 4490, 4492, 4511, 4516, 4517, 4530, 4524, 4522, 4523, 4522, 4531, 4532, 4543, 4549, 4548, 4546, 4537, 4533, 4542, 4551, 4552, 4562, 4570, 4571, 4576, 4585, 4597, 4599, 4600, 4605, 4606, 4607, 4630, 4642, 4643, 4639, 4675, 4661, 4684, 4685, 4697, 4699, 4730, 4731, 4732, 4734, 4737, 4744, 4761, 4771, 4775, 4776, 4784, 4789, 4790, 4767, 4768, 4770, 4771, 4772, 4792, 4801, 4808, 4811, 4818, 4811, 4813, 4814, 4815, 4824, 4823, 4824, 4827, 4830, 4832, 4821, 4809, 4812, 4841, 4842, 4844, 4801, 4805, 4820, 4832, 4823, 4824, 4816, 4818, 4817, 4816, 4822, 4858, 4861, 4880, 4879, 4886, 4893, 4902, 4906, 4901, 4913, 4919, 4928, 4926, 4927, 4935, 4937, 4938, 4939, 4941, 4942, 4947, 4949, 4963, 4955, 4958, 4960, 4956, 4958, 4957, 4961, 4971, 4972, 4975, 4978, 4979, 4985, 4972, 5019, 5023, 5022, 5014, 5016, 5024, 5026, 5049, 5054, 5055, 5051, 5052, 5055, 5062, 5065, 5067, 5065, 5067, 5047, 5060, 5061, 5074, 5101, 5103, 5104, 5108, 5110, 5113, 5124, 5122, 5109, 5110, 5093, 5103, 5099, 5101, 5102, 5108, 5106, 5139, 5160, 5164, 5165, 5170, 5164, 5166, 5168, 5178, 5199, 5200, 5202, 5206, 5214, 5217, 5222, 5223, 5263, 5268, 5288, 5289, 5301, 5324, 5301, 5302, 5304, 5318, 5321, 5329, 5337, 5340, 5346, 5369, 5386, 5377, 5392, 5383, 5385, 5403, 5402, 5404, 5434, 5440, 5442, 5434, 5442, 5452, 5453, 5495, 5499, 5503, 5509, 5508, 5513, 5515, 5521, 5527, 5532, 5543, 5544, 5549, 5550, 5551, 5537, 5508, 5509, 5514, 5517, 5521, 5538, 5546, 5561, 5563, 5566, 5569, 5581, 5592, 5608, 5607, 5612, 5618, 5630, 5636, 5652, 5662, 5663, 5672, 5681, 5682, 5680, 5671, 5678, 5692, 5701, 5707, 5709, 5710, 5711, 5713, 5702, 5713, 5700, 5701, 5702, 5699, 5678, 5679, 5684, 5688, 5720, 5722, 5726, 5731, 5732, 5731, 5735, 5741, 5742, 5747, 5748, 5773, 5792, 5795, 5797, 5820, 5819, 5825, 5827, 5829, 5852, 5853, 5856, 5868, 5875, 5883, 5907, 5872, 5882, 5893, 5929, 5942, 5943, 5946, 5948, 5949, 5950, 5959, 5961, 5965, 5971, 5972, 5973, 5964, 5957, 5977, 5978, 5976, 5980, 5984, 6009, 6023, 6035, 6037, 6039, 6040, 6041, 6048, 6061, 6081, 6080, 6088, 6086, 6089, 6113, 6118, 6124, 6139, 6140, 6142, 6134, 6125, 6130, 6134, 6139, 6145, 6169, 6170, 6173, 6174, 6184, 6155, 6169, 6184, 6190, 6199, 6204, 6205, 6221, 6220, 6229, 6206, 6207, 6209, 6219, 6222, 6220, 6221, 6223, 6225, 6226, 6245, 6246, 6251, 6257, 6258, 6269, 6280, 6282, 6283, 6287, 6296, 6298, 6303, 6304, 6312, 6301, 6304, 6302, 6314, 6317, 6336, 6348, 6357, 6355, 6356, 6364, 6368, 6366, 6374, 6372, 6351, 6358, 6356, 6364, 6365, 6381, 6406, 6420, 6424, 6420, 6432, 6447, 6454, 6456, 6498, 6497, 6498, 6500, 6508, 6544, 6546, 6548, 6557, 6564, 6566, 6567, 6570, 6599, 6602, 6600, 6606, 6608, 6605, 6608, 6610, 6624, 6601, 6604, 6603, 6619, 6647, 6661, 6660, 6663, 6664, 6679, 6682, 6686, 6701, 6702, 6710, 6719, 6720, 6715, 6716, 6718, 6728, 6729, 6712, 6711, 6718, 6730, 6731, 6757, 6760, 6765, 6748, 6747, 6746, 6775, 6776, 6775, 6785, 6786, 6788, 6791, 6792, 6800, 6802, 6803, 6806, 6809, 6806, 6818, 6821, 6823, 6824, 6832, 6835, 6834, 6838, 6840, 6844, 6845, 6854, 6857, 6863, 6885, 6893, 6908, 6922, 6914, 6915, 6922, 6923, 6924, 6926, 6930, 6938, 6970, 6971, 6972, 6971, 6978, 6975, 6978, 6958, 6971, 6977, 6978, 6979, 6982, 6986, 6985, 6988, 7000, 7002, 7007, 7017, 7019, 7040, 7043, 7045, 7046, 7047, 7057, 7061, 7073, 7093, 7095, 7096, 7121, 7129, 7130, 7135, 7139, 7141, 7143, 7145, 7147, 7148, 7149, 7152, 7160, 7170, 7171, 7174, 7175, 7170, 7171, 7198, 7201, 7203, 7204, 7211, 7214, 7215, 7246, 7233, 7238, 7240, 7253, 7254, 7249, 7252, 7256, 7267, 7268, 7274, 7277, 7279, 7285, 7290, 7296, 7298, 7313, 7326, 7332, 7335, 7338, 7345, 7346, 7353, 7356, 7360, 7367, 7392, 7395, 7404, 7416, 7418, 7438, 7451, 7452, 7449, 7453, 7469, 7471, 7473, 7503, 7504, 7517, 7520, 7532, 7533, 7544, 7553, 7565, 7575, 7576, 7575, 7581, 7597, 7633, 7634, 7632, 7633, 7637, 7638, 7644, 7634, 7643, 7644, 7645, 7646, 7649, 7629, 7630, 7632, 7635, 7636, 7642, 7643, 7645, 7644, 7660, 7663, 7672, 7675, 7677, 7680, 7701, 7702, 7709, 7720, 7726, 7734, 7735, 7773, 7774, 7779, 7782, 7787, 7783, 7785, 7787, 7788, 7791, 7793, 7797, 7802, 7803, 7805, 7809, 7805, 7815, 7826, 7832, 7851, 7852, 7853, 7874, 7854, 7858, 7866, 7867, 7850, 7858, 7876, 7892, 7897, 7898, 7900, 7908, 7909, 7911, 7917, 7918, 7919, 7920, 7914, 7911, 7914, 7929, 7942, 7943, 7944, 7945, 7944, 7945, 7976, 7988, 7998, 8001, 8002, 8008, 8010, 8006, 8011, 8006, 8016, 8018, 8019, 8020, 8024, 8027, 8031, 8022, 8025, 8039, 8040, 8049, 8069, 8070, 8071, 8080, 8081, 8093, 8096, 8107, 8110, 8115, 8124, 8126, 8115, 8118, 8121, 8142, 8143, 8147, 8148, 8152, 8156, 8158, 8159, 8160, 8163, 8166, 8179, 8202, 8203, 8205, 8214, 8215, 8211, 8218, 8228, 8229, 8230, 8252, 8255, 8258, 8266, 8267, 8268, 8270, 8275, 8276, 8278, 8283, 8292, 8303, 8305, 8293, 8297, 8313, 8317, 8308, 8307, 8308, 8323, 8326, 8331, 8328, 8329, 8332, 8335, 8337, 8336, 8337, 8338, 8362, 8363, 8364, 8366, 8369, 8379, 8384, 8385, 8392, 8400, 8424, 8429, 8432, 8434, 8433, 8434, 8437, 8438, 8442, 8437, 8438, 8450, 8451, 8445, 8449, 8452, 8478, 8480, 8498, 8490, 8492, 8493, 8494, 8495, 8503, 8504, 8507, 8516, 8517, 8543, 8544, 8555, 8571, 8572, 8574, 8575, 8580, 8590, 8593, 8601, 8603, 8610, 8616, 8617, 8619, 8644, 8645, 8646, 8652, 8653, 8654, 8655, 8656, 8659, 8660, 8666, 8668, 8669, 8672, 8690, 8682, 8685, 8686, 8665, 8674, 8672, 8677, 8678, 8679, 8684, 8695, 8704, 8725, 8745, 8747, 8739, 8755, 8757, 8734, 8738, 8737, 8730, 8735, 8739, 8747, 8750, 8777, 8778, 8780, 8782, 8783, 8785, 8786, 8777, 8778, 8779, 8783, 8788, 8789, 8795, 8811, 8829, 8834, 8862, 8871, 8875, 8878, 8879, 8872, 8867, 8881, 8892, 8893, 8896, 8897, 8896, 8897, 8898, 8907, 8908, 8910, 8937, 8943, 8950, 8951, 8953, 8964, 8985, 8982, 8987, 8994, 9021, 9023, 9024, 9019, 9024, 9029, 9008, 9020, 9026, 9027, 9044, 9061, 9058, 9061, 9065, 9066, 9067, 9077, 9078, 9077, 9075, 9076, 9078, 9087, 9082, 9083, 9085, 9086, 9101, 9112, 9122, 9123, 9128, 9129, 9131, 9134, 9138, 9147, 9153, 9171, 9165, 9155, 9159, 9185, 9183, 9186, 9190, 9193, 9194, 9197, 9203, 9204, 9214, 9213, 9218, 9227, 9231, 9247, 9262, 9266, 9268, 9269, 9259, 9260, 9268, 9278, 9279, 9299, 9300, 9301, 9302, 9312, 9313, 9314, 9346, 9347, 9348, 9337, 9342, 9338, 9337, 9357, 9360, 9364, 9367, 9382, 9410, 9411, 9422, 9429, 9431, 9430, 9437, 9438, 9442, 9441, 9443, 9455, 9462, 9464, 9472, 9497, 9528, 9551, 9553, 9555, 9556, 9569, 9580, 9584, 9586, 9575, 9578, 9579, 9587, 9590, 9619, 9622, 9628, 9616, 9614, 9625, 9626, 9631, 9641, 9648, 9658, 9663, 9664, 9665, 9649, 9650, 9647, 9657, 9661, 9663, 9668, 9673, 9678, 9681, 9682, 9687, 9688, 9705, 9706, 9710, 9716, 9728, 9730, 9733, 9726, 9727, 9728, 9729, 9730, 9727, 9729, 9731, 9732, 9735, 9746, 9748, 9753, 9752, 9739, 9738, 9760, 9765, 9768, 9767, 9771, 9762, 9765, 9766, 9759, 9765, 9766, 9797, 9802, 9803, 9805, 9807, 9813, 9814, 9818, 9819, 9825, 9826, 9827, 9828, 9829, 9833, 9849, 9855, 9856, 9875, 9893, 9899}
}

type DayOneTest struct {
	input    []uint
	expected uint
}

func TestDayOne(t *testing.T) {
	tests := []DayOneTest{
		{
			[]uint{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			7,
		},
		{
			[]uint{1, 2, 1, 10, 1},
			2,
		},
		{
			[]uint{1, 1, 1, 1, 1},
			0,
		},
	}
	for _, test := range tests {
		actual := DayOne(test.input)
		if actual != test.expected {
			t.Fatalf(`Error! Output: %v Test: %v`, actual, test)
		}
	}
}
