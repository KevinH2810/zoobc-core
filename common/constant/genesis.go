// Constants to generate/validate genesis block
// Note: this file has been auto-generated by 'genesis generate' command
package constant

const (
	MainchainGenesisBlockID int64 = 2447830660844865454
)

type (
	GenesisConfigEntry struct {
		AccountAddress     string
		AccountBalance     int64
		NodePublicKey      []byte
		NodeAddress        string
		LockedBalance      int64
		ParticipationScore int64
	}
)

var (
	MainchainGenesisBlocksmithID   = []byte{1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1, 1}
	MainchainGenesisBlockSignature = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	MainchainGenesisTransactionSignature = []byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
		0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	MainchainGenesisBlockTimestamp = int64(1596790800)
	MainchainGenesisAccountAddress = "ZBC_AQTEH7K4_L45WJPLL_HCEC65ZH_7XC5N3XD_YNKPHK45_POH7PQME_AFAFBDWM"
	MainchainGenesisBlockSeed      = make([]byte, 64)
	MainchainGenesisNodePublicKey  = make([]byte, 32)
	GenesisConfig                  = []GenesisConfigEntry{
		{
			AccountAddress: "ZBC_4WN2TZMB_J2AVKGI5_DNUM7AJ5_7PBKFWPT_ORJ4BUY4_EM5FWFXP_QBUM6D6Z",
			AccountBalance: 0,
			NodePublicKey: []byte{8, 148, 20, 54, 230, 172, 28, 39, 200, 151, 101, 206, 38, 100, 140, 223, 39, 128, 72,
				132, 78, 165, 61, 63, 185, 144, 98, 228, 192, 25, 77, 66},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_XDWDR6YZ_7XMB7MXA_5ZGYTBDL_ZZKEA7OE_TGFXJXQB_IHJY42PG_CYC4OED7",
			AccountBalance: 0,
			NodePublicKey: []byte{121, 79, 159, 56, 18, 2, 59, 190, 133, 35, 233, 41, 99, 228, 129, 175, 162, 196, 194,
				24, 134, 180, 239, 188, 221, 235, 221, 104, 40, 132, 175, 63},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_6RTZQVVU_SLT6AKNX_SCWQTCEZ_6EKW2HU2_YQ3IIX3U_BCKTW3LX_HRRQ3CZE",
			AccountBalance: 0,
			NodePublicKey: []byte{146, 72, 54, 42, 91, 39, 3, 191, 27, 86, 59, 97, 37, 84, 240, 17, 235, 246, 88,
				162, 165, 171, 137, 173, 94, 192, 192, 251, 96, 243, 98, 219},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_LU4QSNY3_G6VC3G4N_YB2IYKJA_UNSJMTAI_TWBGOQW7_ZTT6DW5A_BIZAE42I",
			AccountBalance: 0,
			NodePublicKey: []byte{212, 188, 117, 21, 238, 133, 48, 111, 10, 162, 113, 11, 223, 51, 101, 27, 42, 109, 223,
				2, 162, 72, 129, 180, 122, 8, 120, 185, 222, 128, 236, 221},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_XDGAN652_ZLS4LN7Y_NNBU2KDH_LM2KDD2E_NGUIKTEP_E72CHSHT_ISLOSO5A",
			AccountBalance: 0,
			NodePublicKey: []byte{142, 56, 183, 28, 8, 48, 154, 86, 37, 13, 171, 105, 181, 58, 10, 97, 2, 172, 90,
				223, 148, 150, 30, 88, 17, 186, 121, 228, 123, 40, 22, 70},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_46JBB3TX_QL7OAUIB_CRKSOGKU_Q73HD2D4_M7IS6YMN_SH5J3UWU_V3OOLH3K",
			AccountBalance: 0,
			NodePublicKey: []byte{212, 0, 142, 14, 160, 71, 143, 200, 41, 225, 211, 29, 46, 178, 219, 182, 153, 143, 254,
				103, 35, 36, 236, 167, 133, 238, 194, 178, 226, 209, 28, 190},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_WN7ATGN7_JTSN6OSH_R3UQFRSK_K5BVKK6O_7EDZ24QN_S4E2L5S7_4KF76EIF",
			AccountBalance: 0,
			NodePublicKey: []byte{55, 182, 221, 101, 151, 72, 133, 5, 128, 89, 128, 15, 41, 209, 242, 18, 28, 82, 91,
				182, 55, 247, 116, 50, 209, 185, 17, 137, 30, 17, 63, 18},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_7HHAYQI3_DQMIFPOW_OPZKCYQY_X5GFDTX6_QR7OIA2C_JPX6HIW5_AKNRSYTH",
			AccountBalance: 0,
			NodePublicKey: []byte{240, 223, 0, 74, 104, 19, 226, 255, 239, 88, 82, 196, 54, 183, 138, 162, 136, 22, 172,
				26, 44, 58, 87, 7, 5, 153, 137, 120, 231, 198, 218, 51},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_WQCF6U5S_GFHLDTMZ_36OD6GB3_4JFIRWHB_5Q37B7X6_LEBNM4UY_Q2FM5EI7",
			AccountBalance: 0,
			NodePublicKey: []byte{186, 113, 9, 153, 55, 189, 46, 128, 175, 24, 119, 229, 202, 31, 12, 164, 86, 255, 40,
				198, 28, 86, 221, 104, 233, 117, 153, 88, 197, 251, 190, 27},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_Z2BM2FMO_HQ2UDPJX_N6FGA6UF_APP6YLJH_7GLFX7PL_X3O3THJV_HPHN4PMK",
			AccountBalance: 0,
			NodePublicKey: []byte{170, 200, 220, 151, 50, 209, 189, 245, 164, 162, 189, 61, 81, 37, 226, 210, 203, 218, 129,
				76, 233, 168, 66, 56, 81, 176, 114, 27, 142, 73, 252, 17},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_A7EOF4R3_PIPT7LNO_B6UDMN24_XOVFDMUD_ECAOVHRT_SB4KYEDI_3ES6FGV2",
			AccountBalance: 0,
			NodePublicKey: []byte{222, 244, 51, 142, 143, 110, 116, 205, 238, 237, 195, 121, 203, 157, 122, 253, 173, 25, 35,
				145, 39, 151, 15, 152, 81, 239, 74, 71, 205, 71, 198, 129},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_PPFNBPUX_OXULYAMQ_HDFAZVW4_URXYOUIF_HNAJ2RGK_7MYD476M_WOI7QPUK",
			AccountBalance: 0,
			NodePublicKey: []byte{144, 123, 254, 132, 236, 165, 225, 64, 121, 115, 79, 47, 245, 139, 0, 149, 189, 172, 181,
				134, 36, 195, 204, 201, 9, 167, 55, 175, 39, 71, 4, 120},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_AHCI6OI3_I7V5W4TI_ZZBYIGAP_DQ4XHNMA_BUROUOLX_76QMLOTL_IYC6A6UU",
			AccountBalance: 0,
			NodePublicKey: []byte{154, 176, 57, 36, 219, 114, 64, 130, 219, 160, 118, 212, 182, 187, 79, 244, 114, 25, 106,
				55, 152, 22, 63, 187, 205, 117, 10, 56, 129, 252, 158, 235},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_2OOVGEXU_TZ66FEFY_KECXOMRB_RNKPEKWD_IYN463L5_NTMBXID3_G55L36NN",
			AccountBalance: 0,
			NodePublicKey: []byte{209, 133, 193, 219, 162, 249, 16, 63, 148, 232, 9, 164, 131, 110, 118, 221, 12, 184, 197,
				46, 95, 152, 59, 235, 31, 195, 123, 133, 67, 17, 2, 93},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_DJXBF2O2_CRQDVPSK_QYZEFYAY_DIGRLE2K_HT6NZTWA_7V4SINRK_XJYJ337I",
			AccountBalance: 0,
			NodePublicKey: []byte{145, 198, 179, 228, 55, 4, 237, 141, 201, 210, 252, 69, 47, 102, 27, 145, 233, 230, 93,
				86, 175, 88, 48, 248, 84, 128, 182, 209, 248, 213, 77, 72},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_KLM5K5WH_IX7K2XC3_I22QRHF3_EVMJY7AS_UJ3TVL5E_GUUCDPG5_FH2BLYPJ",
			AccountBalance: 0,
			NodePublicKey: []byte{245, 237, 170, 105, 47, 139, 45, 121, 185, 62, 50, 55, 191, 116, 164, 196, 108, 183, 216,
				221, 201, 61, 62, 88, 92, 105, 194, 104, 60, 79, 57, 22},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_DFGEYAOA_DHMAJIGW_V3NY32BK_GVMWLAKY_6XSW4LW3_XJOU232C_3ZF6UHBA",
			AccountBalance: 0,
			NodePublicKey: []byte{98, 207, 144, 113, 120, 170, 253, 35, 191, 66, 0, 64, 19, 231, 241, 60, 53, 54, 74,
				194, 104, 129, 121, 17, 53, 31, 46, 58, 133, 127, 225, 128},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_ONELFEGT_NLNLFIUE_G43U2BBN_RVTRCRUU_OHMAAGYL_24L2U5SN_Q4BV4CLZ",
			AccountBalance: 0,
			NodePublicKey: []byte{179, 203, 71, 144, 150, 224, 6, 200, 172, 78, 194, 211, 95, 79, 65, 234, 50, 105, 76,
				36, 191, 228, 6, 127, 211, 236, 46, 199, 169, 76, 29, 3},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_YG7S4ZQD_QNJD3NYB_L3DAXVKE_JXVKSW7Z_JPBGIUC2_JIPUST2P_WZRW7WWT",
			AccountBalance: 0,
			NodePublicKey: []byte{96, 202, 238, 114, 232, 44, 66, 217, 160, 144, 139, 31, 43, 13, 20, 42, 208, 89, 76,
				83, 138, 11, 182, 179, 47, 232, 220, 234, 244, 230, 209, 41},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_VY36OGW3_XFGAIVDQ_F5GDXPDG_2NYHKISU_6NUYPTMA_27D4JIO7_BZ5KRKXH",
			AccountBalance: 0,
			NodePublicKey: []byte{6, 19, 172, 83, 235, 130, 230, 89, 30, 94, 196, 165, 177, 249, 186, 130, 153, 236, 137,
				54, 165, 208, 79, 204, 115, 137, 9, 197, 14, 142, 209, 241},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_EVFAZ66R_SYKPLU75_6ZRL5NC6_ZAB2OXIQ_A2QS67VE_HUNMZ2NT_TJMXH72E",
			AccountBalance: 0,
			NodePublicKey: []byte{57, 24, 237, 57, 154, 172, 75, 152, 21, 104, 145, 108, 131, 89, 92, 31, 76, 18, 227,
				239, 187, 220, 229, 208, 7, 102, 70, 31, 92, 1, 52, 244},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_OBUTEVO4_OI6LZDLQ_FBKXOQJO_LWGJKRN7_PBISIYKC_2WBISBXN_QRJNKP6X",
			AccountBalance: 0,
			NodePublicKey: []byte{249, 157, 167, 69, 234, 160, 74, 116, 12, 66, 132, 109, 106, 53, 127, 120, 117, 62, 65,
				221, 114, 228, 247, 70, 201, 159, 102, 76, 158, 43, 215, 23},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_6TGXQBQF_F3FCW4ZV_BQLAOPLT_66UVQC4P_73MJ6RUY_EDHXTZRI_KBB2ZJVO",
			AccountBalance: 0,
			NodePublicKey: []byte{9, 82, 45, 72, 139, 59, 173, 1, 197, 211, 1, 98, 192, 109, 138, 253, 53, 24, 79,
				160, 76, 240, 229, 43, 211, 201, 19, 164, 121, 43, 116, 54},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_WP5FMEVF_E5FXR4YT_WPBD63HU_K44MRCIY_UO267C7D_TC7OSOL7_TU7Y2N34",
			AccountBalance: 0,
			NodePublicKey: []byte{248, 162, 107, 7, 244, 249, 149, 132, 22, 231, 244, 19, 171, 148, 225, 44, 153, 214, 207,
				37, 82, 41, 8, 225, 115, 8, 153, 185, 246, 141, 118, 71},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_PUGFNG5U_EEU2VTA7_RMSVD6K6_D27U6YFH_UJNRZUW3_5UWROAUY_OHPW5YO2",
			AccountBalance: 0,
			NodePublicKey: []byte{54, 23, 3, 199, 36, 141, 159, 9, 94, 73, 0, 143, 17, 77, 35, 142, 152, 34, 92,
				176, 17, 120, 172, 37, 14, 14, 58, 124, 229, 154, 169, 25},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_5JI2UT5O_RGEBHNDQ_5NA4MN5D_CCEMMOFZ_LKECWJFO_U7GXDRNG_SQGFTGAH",
			AccountBalance: 0,
			NodePublicKey: []byte{113, 63, 39, 31, 202, 235, 82, 246, 112, 68, 50, 230, 15, 23, 24, 255, 27, 103, 187,
				143, 194, 173, 149, 87, 59, 162, 137, 28, 96, 158, 246, 136},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_JA2HEKPV_2RQ5CLPU_KHBWNWY2_KJLQMCRH_VHS4ZN3U_XSXXMONS_GPFHSZ46",
			AccountBalance: 0,
			NodePublicKey: []byte{154, 34, 217, 141, 174, 202, 214, 173, 25, 52, 130, 184, 146, 223, 51, 205, 145, 167, 136,
				228, 12, 203, 99, 83, 52, 255, 50, 64, 200, 103, 57, 0},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_JCCHYZGN_CMZMGTLY_43SJCR2C_6EUFPDMJ_CHIIZINY_7DKHWMQW_YMQQLCB3",
			AccountBalance: 0,
			NodePublicKey: []byte{238, 51, 72, 233, 40, 241, 237, 31, 91, 237, 168, 109, 187, 217, 209, 182, 38, 100, 44,
				100, 14, 251, 51, 218, 135, 56, 216, 114, 91, 18, 232, 20},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_JGJCIBVS_XCFLT3KL_46TUAEDO_62AZEKPH_3J4KQQUE_77M3DJBC_GJEYYN3Q",
			AccountBalance: 0,
			NodePublicKey: []byte{227, 100, 78, 99, 33, 31, 203, 224, 155, 187, 51, 65, 58, 142, 79, 179, 106, 192, 90,
				191, 172, 98, 87, 0, 44, 198, 10, 14, 100, 5, 193, 232},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_5J3ALALF_AM5KG2T2_LLWVSI5O_5F4KX226_ZWSICFHN_HS362D2K_KPCZHWOJ",
			AccountBalance: 0,
			NodePublicKey: []byte{68, 6, 12, 125, 209, 135, 114, 126, 224, 96, 98, 55, 44, 221, 206, 242, 162, 170, 24,
				33, 76, 210, 173, 12, 40, 111, 100, 244, 195, 75, 187, 251},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_FDYMDD5P_I3ORLYVA_NFKP7WST_5PLJQGSF_JWEB62IX_RSTZJQYF_XA32OJWL",
			AccountBalance: 0,
			NodePublicKey: []byte{195, 57, 163, 138, 144, 225, 164, 27, 180, 37, 44, 82, 235, 148, 167, 238, 29, 47, 123,
				117, 225, 203, 40, 232, 42, 15, 38, 41, 140, 55, 144, 13},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_5LPFWQFS_VQ4ZCDSO_674RJGS6_UDZ5QIG4_DC6COKKH_VQE4JJF6_NZFSNG7U",
			AccountBalance: 0,
			NodePublicKey: []byte{188, 133, 14, 62, 246, 191, 92, 240, 70, 30, 154, 122, 230, 191, 251, 185, 225, 214, 24,
				169, 127, 7, 20, 54, 179, 252, 103, 21, 220, 61, 202, 211},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_NVZILSTG_YKQDUUXZ_V7QZEYIL_QIGAO4P7_DTUPST2E_4YNKEMWF_CFCKMFY6",
			AccountBalance: 0,
			NodePublicKey: []byte{69, 255, 141, 96, 94, 192, 141, 209, 188, 140, 168, 235, 213, 11, 44, 180, 210, 227, 240,
				243, 198, 34, 71, 222, 186, 207, 62, 99, 12, 175, 31, 217},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_GUYNIKXT_LCI2J4TC_44JADBJN_IJLEOGWE_UAGH5LSA_R2NNAQCU_THYT7SNM",
			AccountBalance: 0,
			NodePublicKey: []byte{176, 186, 50, 116, 37, 201, 69, 89, 200, 49, 190, 240, 40, 246, 92, 58, 141, 56, 149,
				239, 199, 171, 123, 96, 96, 120, 75, 191, 142, 66, 23, 163},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_XZW2UMFN_7WGUYU5T_GVNXPQRM_BF4M2X4A_JUXSMWS3_K237NPOV_6XUL22HA",
			AccountBalance: 0,
			NodePublicKey: []byte{59, 140, 215, 104, 172, 228, 117, 196, 82, 183, 71, 97, 231, 201, 157, 179, 80, 202, 9,
				66, 38, 116, 140, 200, 98, 71, 205, 155, 198, 228, 213, 173},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_3HOWUZ64_ZY3KHADN_4MOMV2RK_TTJZFHWJ_FTNQ5T7O_NBYSLIXM_G55GYW53",
			AccountBalance: 0,
			NodePublicKey: []byte{147, 199, 230, 172, 179, 86, 102, 95, 236, 110, 252, 192, 3, 0, 23, 155, 140, 62, 1,
				224, 223, 139, 230, 77, 233, 142, 170, 24, 3, 159, 80, 202},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_YKQCW2LF_HW53DZNQ_Z3OSA4ZZ_SJR5CJR7_SYMMFHCA_QPUGGGYN_7MBJ222Q",
			AccountBalance: 0,
			NodePublicKey: []byte{167, 35, 245, 136, 30, 150, 165, 209, 24, 206, 60, 52, 164, 61, 251, 251, 162, 6, 62,
				113, 132, 32, 235, 167, 11, 139, 122, 40, 69, 11, 213, 149},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_2PVTGDMB_LYJZ4KT3_OILS46LV_OBMOKUD7_PRYNBRKB_2OGUC6OQ_7A5B6KCI",
			AccountBalance: 0,
			NodePublicKey: []byte{182, 36, 217, 90, 146, 10, 146, 205, 227, 103, 243, 132, 186, 57, 200, 229, 59, 23, 25,
				119, 2, 51, 112, 187, 169, 214, 239, 192, 202, 6, 126, 39},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_IR4HBPJ3_WR756CXI_YT6SP3TZ_U3PMZNVH_PYDQTYF7_XSY7DRX6_APSRME5I",
			AccountBalance: 0,
			NodePublicKey: []byte{239, 5, 228, 215, 185, 75, 242, 132, 137, 69, 155, 36, 225, 74, 108, 163, 45, 176, 170,
				251, 82, 0, 50, 173, 124, 75, 135, 62, 18, 160, 105, 121},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_HWX2UF2Y_EHUL42MB_GVTNNFHO_I7NDTDVY_YOPJOQO4_WDWOLNIR_U7Z6JU5X",
			AccountBalance: 0,
			NodePublicKey: []byte{218, 163, 169, 173, 56, 69, 82, 212, 196, 6, 89, 83, 83, 88, 22, 240, 83, 192, 138,
				220, 67, 71, 213, 37, 93, 150, 74, 158, 164, 107, 221, 12},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_L2UTU3VN_AQE7BYFP_RRI35PGT_6XZ5XE4F_A3KCUPUA_7Q5BXZZO_QG5X4QYT",
			AccountBalance: 0,
			NodePublicKey: []byte{51, 188, 219, 211, 71, 228, 248, 225, 82, 84, 99, 231, 184, 208, 41, 217, 166, 110, 121,
				241, 181, 251, 224, 160, 199, 175, 196, 198, 179, 171, 246, 182},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_DFBAGLCD_64FYCWC3_GUP73NO3_GTXMXJAD_XXXEUFIO_DHZCK7N5_ZFETNYUV",
			AccountBalance: 0,
			NodePublicKey: []byte{190, 37, 235, 97, 72, 15, 137, 113, 11, 27, 187, 74, 185, 58, 148, 179, 10, 164, 26,
				167, 101, 52, 152, 34, 245, 66, 169, 217, 146, 241, 28, 200},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_57BR3N2G_J3P6LBLH_QJTSSWPU_4SFSHGBJ_Y7HLXLC7_ARXALY4S_7F2UGV6W",
			AccountBalance: 0,
			NodePublicKey: []byte{228, 173, 169, 152, 87, 131, 120, 131, 189, 61, 254, 45, 78, 191, 226, 58, 248, 207, 104,
				219, 187, 5, 163, 135, 59, 158, 234, 5, 204, 2, 146, 26},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_B4LQLBYY_7FXJ34QL_KC43NH25_F5A53ZHU_4LRZ32G5_GDBQCVAJ_7AQLLWTK",
			AccountBalance: 0,
			NodePublicKey: []byte{207, 65, 135, 81, 87, 237, 71, 207, 127, 14, 107, 134, 128, 72, 0, 201, 215, 45, 249,
				204, 235, 248, 140, 220, 67, 151, 106, 132, 45, 194, 41, 23},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_FCEU3CCJ_SUZWNLKD_QZDNGRJA_MJ25HMOE_WFUKHLB4_QZOYTPNT_TOEFS3XL",
			AccountBalance: 0,
			NodePublicKey: []byte{55, 16, 123, 251, 161, 114, 67, 99, 86, 48, 232, 23, 182, 70, 52, 8, 215, 34, 49,
				95, 59, 214, 121, 23, 107, 112, 145, 58, 137, 119, 96, 120},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_G7Y6AFOK_B4G6VQV7_NYZB7L4B_DOYANSMD_KUE3JJW7_37MDK64G_VVBL4SNN",
			AccountBalance: 0,
			NodePublicKey: []byte{114, 253, 8, 95, 184, 138, 13, 209, 71, 207, 254, 183, 161, 89, 37, 137, 86, 253, 102,
				28, 25, 237, 38, 255, 227, 76, 165, 6, 130, 29, 184, 229},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_63524TBN_6GYDONEN_5LPOLXOA_FJXD7EUI_DAQQZFCN_N6MN54WO_TE5E64YK",
			AccountBalance: 0,
			NodePublicKey: []byte{36, 120, 136, 64, 208, 158, 15, 98, 76, 219, 127, 106, 115, 121, 47, 227, 116, 94, 14,
				32, 18, 156, 217, 177, 25, 10, 43, 52, 19, 4, 62, 203},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_WIFYS2D3_C27F632Q_KDM5J4PC_KELKTXZH_3BDMFSJX_VW4CGCR4_RNDWNP46",
			AccountBalance: 0,
			NodePublicKey: []byte{36, 115, 69, 104, 215, 248, 34, 227, 247, 27, 28, 125, 38, 192, 52, 6, 185, 178, 160,
				223, 154, 80, 115, 47, 124, 19, 209, 198, 254, 14, 100, 135},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_KT6F37M5_E5HFLYJJ_5W3L7MZE_JMKGFVYP_PNMIZ7NE_ABQC4RA2_GVE547HI",
			AccountBalance: 0,
			NodePublicKey: []byte{141, 224, 150, 110, 70, 255, 130, 56, 194, 121, 134, 35, 53, 12, 7, 47, 180, 53, 84,
				126, 159, 82, 115, 48, 112, 230, 186, 62, 177, 88, 30, 162},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_N6VA4NAR_GDU77F3K_MTUWWMKW_D4EB7LB4_QGXHTZHQ_5S55LVOP_P4L2OISI",
			AccountBalance: 0,
			NodePublicKey: []byte{33, 221, 55, 189, 114, 57, 160, 5, 62, 242, 89, 94, 127, 161, 9, 250, 231, 13, 96,
				220, 62, 185, 86, 126, 134, 209, 240, 96, 227, 94, 228, 171},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_RPGBPNL4_WIYTZFTO_XTT2KOFX_QA4ESTS3_IYBOTYHW_65PNVH3D_67NVXPPZ",
			AccountBalance: 0,
			NodePublicKey: []byte{125, 183, 233, 190, 189, 5, 145, 206, 172, 20, 157, 188, 249, 89, 206, 191, 4, 139, 183,
				65, 70, 136, 127, 103, 145, 196, 42, 167, 82, 137, 189, 223},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_QMTOFZ3Q_7ABTKLS4_CRKGUNQM_SB3VMIL6_5MQEY3GV_HRWWMWID_P2W3DS4X",
			AccountBalance: 0,
			NodePublicKey: []byte{42, 214, 63, 142, 173, 149, 130, 185, 238, 114, 192, 89, 142, 100, 230, 196, 35, 28, 91,
				200, 7, 128, 65, 81, 124, 84, 23, 167, 108, 245, 140, 24},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_3P27EBCJ_VJHDIP24_O6ELXAA4_KBBJKFVA_N6WZ66EA_URRB2JN2_2DFFVUFN",
			AccountBalance: 0,
			NodePublicKey: []byte{63, 44, 54, 81, 93, 207, 235, 124, 80, 226, 200, 196, 62, 68, 213, 30, 108, 87, 21,
				181, 113, 215, 66, 122, 39, 210, 115, 38, 71, 158, 92, 60},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_QLVY6ZY4_73BYCSP5_VMDEBY5M_YGHPK5I2_AU32CC3H_AWET722N_6UDMDQTC",
			AccountBalance: 0,
			NodePublicKey: []byte{51, 141, 137, 5, 62, 42, 144, 162, 253, 154, 118, 55, 175, 160, 246, 198, 248, 99, 200,
				195, 106, 130, 135, 29, 115, 35, 111, 192, 27, 182, 160, 71},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_EGNLEHRN_TDATPL6K_EUIKNNGB_HQFF66KU_IN7CDNI6_HZCEE3XE_5AYAQIPS",
			AccountBalance: 0,
			NodePublicKey: []byte{198, 0, 80, 173, 1, 182, 75, 171, 118, 150, 131, 4, 241, 166, 142, 246, 236, 207, 210,
				34, 83, 205, 158, 2, 98, 106, 238, 253, 251, 72, 98, 33},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_JDZ5PWCC_2HW54GGH_FDBBJOXY_UF2GLQMN_Z5KDVINK_CZSTLCLJ_3BWPMPVH",
			AccountBalance: 0,
			NodePublicKey: []byte{186, 100, 212, 82, 157, 137, 209, 228, 99, 105, 81, 22, 208, 93, 125, 236, 3, 55, 100,
				114, 81, 140, 184, 156, 181, 50, 114, 8, 250, 78, 129, 237},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_FCWKWE5S_G7VLNCQZ_HF6VAGZU_BA3YNNUA_5JN4JTUS_2FMQUB7A_DKUNYHHU",
			AccountBalance: 0,
			NodePublicKey: []byte{0, 154, 155, 85, 46, 22, 28, 121, 188, 70, 151, 11, 111, 220, 94, 42, 201, 16, 67,
				8, 240, 234, 214, 196, 158, 72, 63, 20, 240, 61, 57, 187},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_Q7ESAVBY_Y5MP5B3G_AO7UUELL_RIKK2M3I_XGA7K3TX_5H7ZXYGA_UINORFSV",
			AccountBalance: 0,
			NodePublicKey: []byte{217, 146, 87, 155, 142, 255, 212, 12, 122, 224, 216, 180, 254, 33, 121, 236, 21, 113, 231,
				249, 19, 12, 168, 143, 43, 168, 103, 73, 98, 87, 75, 208},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_QKSCKGXT_JR5F77BW_AO6FZBME_LZVONDJL_XDKK6DCC_4OHVBBGW_C46L5UEK",
			AccountBalance: 0,
			NodePublicKey: []byte{84, 193, 97, 123, 73, 11, 74, 63, 17, 224, 77, 192, 140, 43, 144, 20, 150, 11, 188,
				229, 170, 87, 240, 215, 32, 35, 130, 210, 28, 5, 188, 76},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_OI7PBPLS_L6J3UGUY_KNTSJE65_LJKQ2XYP_PKRAXDEM_3Q3MQZZ7_FG5TXICL",
			AccountBalance: 0,
			NodePublicKey: []byte{86, 185, 234, 1, 214, 221, 153, 28, 39, 193, 143, 192, 129, 193, 59, 231, 33, 79, 220,
				80, 59, 92, 246, 77, 107, 93, 200, 136, 162, 199, 169, 217},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_5CRBNZ4O_P45I7UYZ_23QZE5IC_CV36DHJF_J525DEBN_4COXFO5Q_QE35TTY4",
			AccountBalance: 0,
			NodePublicKey: []byte{36, 242, 126, 223, 144, 94, 115, 132, 80, 168, 183, 1, 64, 229, 192, 20, 147, 37, 193,
				202, 146, 122, 6, 44, 254, 136, 95, 116, 25, 64, 159, 174},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_JOJVUA6J_VARDFRUU_6RDFA6XX_R5A34JSG_GNHTYK4E_GHYYXC43_FPF35743",
			AccountBalance: 0,
			NodePublicKey: []byte{235, 226, 22, 2, 229, 34, 58, 158, 169, 166, 185, 194, 138, 32, 253, 117, 145, 183, 105,
				113, 43, 52, 189, 126, 57, 102, 110, 33, 226, 202, 0, 224},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_CFABQOIF_GW2YVOK6_AA2ILXAP_NRQ253EH_ZV5SB7IN_D3TTFB32_WECVCYPF",
			AccountBalance: 0,
			NodePublicKey: []byte{52, 203, 116, 184, 91, 125, 118, 207, 139, 203, 15, 205, 73, 246, 66, 132, 181, 199, 6,
				200, 77, 118, 195, 195, 140, 65, 218, 142, 211, 110, 102, 40},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_TST4NG5Q_O5E7A3DO_R3SLEKXK_RAJQQQZO_OFBQWP3O_XLN7TSZS_Q5KKF5OG",
			AccountBalance: 0,
			NodePublicKey: []byte{134, 192, 144, 100, 150, 223, 114, 191, 234, 60, 13, 223, 103, 104, 218, 74, 52, 1, 11,
				161, 61, 143, 183, 78, 245, 157, 103, 231, 15, 196, 127, 245},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_HUUAVTD5_YGZJNFOS_D5TAB7JP_XC3PMLRF_GH6X724D_Q7R52MUF_K7V3XSBW",
			AccountBalance: 0,
			NodePublicKey: []byte{82, 45, 126, 118, 77, 109, 229, 222, 176, 22, 121, 184, 194, 54, 168, 71, 78, 208, 84,
				27, 220, 193, 132, 193, 39, 10, 60, 152, 146, 160, 223, 148},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_6HJSGRVL_Q5HVVNHI_N2U4AR43_DAF23XXI_XTYQUK6H_OF4MC5Y6_EVD4U3WO",
			AccountBalance: 0,
			NodePublicKey: []byte{184, 125, 82, 121, 145, 199, 72, 86, 233, 171, 94, 13, 140, 178, 174, 212, 119, 9, 110,
				103, 41, 21, 175, 128, 50, 28, 61, 241, 192, 67, 243, 207},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_Z2NXJFQN_7DXLZHSL_5PAWHK5J_OKSTEO3V_JPA5WCTI_K2RUT7QL_3G2OFRS4",
			AccountBalance: 0,
			NodePublicKey: []byte{221, 90, 77, 123, 69, 136, 5, 173, 93, 183, 91, 9, 128, 36, 193, 128, 225, 101, 136,
				51, 33, 162, 218, 41, 59, 219, 139, 248, 143, 244, 86, 111},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_KQT3BUCP_AARLPL3J_Y7TOAQJV_NLPHELX3_E2VR3HCL_A5E6M5AL_QJVRTS26",
			AccountBalance: 0,
			NodePublicKey: []byte{36, 13, 31, 191, 72, 105, 27, 48, 113, 29, 190, 151, 35, 76, 107, 87, 221, 148, 118,
				115, 73, 199, 232, 128, 87, 112, 198, 217, 159, 161, 180, 50},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_H7BDNKZ5_43WVRVXR_ZN6XDJF4_67BSLVXY_XCA3KH4A_ARAAF6DJ_2FHFPH2X",
			AccountBalance: 0,
			NodePublicKey: []byte{198, 210, 38, 60, 255, 112, 121, 52, 27, 163, 195, 163, 88, 187, 94, 208, 142, 209, 223,
				141, 46, 55, 110, 179, 167, 237, 19, 25, 195, 142, 44, 109},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_YI7ZEPEQ_RAFG2JOK_IDQVI4T7_RMGSBRSA_GTZI646V_CF5LGDNT_5FEPDMUO",
			AccountBalance: 0,
			NodePublicKey: []byte{37, 147, 157, 4, 197, 241, 176, 12, 247, 77, 239, 238, 198, 86, 43, 26, 180, 25, 240,
				155, 25, 62, 113, 43, 175, 117, 5, 43, 236, 42, 169, 126},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_JTAFVCSO_M2OPFHBV_2K4VGHVT_NGAYYV4A_ZB3JLFJZ_SWCK4YBN_2UTKE3WB",
			AccountBalance: 0,
			NodePublicKey: []byte{19, 34, 134, 63, 213, 178, 131, 18, 29, 174, 125, 3, 111, 178, 100, 28, 233, 5, 141,
				90, 140, 196, 11, 99, 166, 30, 123, 191, 254, 236, 163, 103},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_YIFR27XU_EONFCPN6_5APOGACF_D26IOT2B_CNIFB3JM_5IYCPUHR_56TJT2W7",
			AccountBalance: 0,
			NodePublicKey: []byte{17, 93, 127, 35, 39, 240, 182, 245, 51, 39, 239, 98, 16, 28, 210, 104, 228, 84, 30,
				251, 152, 221, 209, 29, 10, 254, 85, 70, 191, 201, 35, 42},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_44R2V2TL_NSSHAQFG_RQ3HJ4RI_KHISLCRD_ODAGWDX4_5IGV6ABT_4ZVZP3ZK",
			AccountBalance: 0,
			NodePublicKey: []byte{127, 251, 51, 131, 68, 107, 26, 38, 227, 219, 130, 243, 214, 245, 224, 5, 250, 143, 149,
				55, 42, 45, 252, 54, 163, 152, 20, 84, 13, 152, 162, 28},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_ZXZSX3GW_CYXZYX3E_UEW4LSJV_TSCJVNCU_2WCUSN34_FCQDO72M_PHFLVB4Y",
			AccountBalance: 0,
			NodePublicKey: []byte{228, 77, 56, 134, 135, 184, 36, 134, 41, 158, 244, 147, 48, 203, 197, 175, 207, 134, 24,
				238, 197, 236, 4, 63, 177, 36, 79, 201, 38, 160, 83, 19},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_DVGEW4JU_4L6A7NNK_67A7PBBL_JMKBOTNE_KY64KDXC_7XQGQBCX_OHLUPZNZ",
			AccountBalance: 0,
			NodePublicKey: []byte{158, 19, 115, 37, 200, 113, 176, 59, 136, 116, 159, 82, 82, 163, 22, 92, 122, 12, 105,
				185, 34, 54, 241, 143, 74, 140, 200, 254, 126, 191, 211, 199},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_LZNHMUQ3_BKYP5574_5GA33DEG_HWRIQ67D_UFSG5DCP_ZLL6YIET_DFGZD6G6",
			AccountBalance: 0,
			NodePublicKey: []byte{15, 159, 25, 77, 136, 147, 134, 193, 170, 32, 46, 38, 145, 11, 253, 146, 104, 118, 143,
				239, 95, 186, 177, 137, 152, 79, 243, 47, 130, 196, 127, 221},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_CWOEY766_5UK3ZVAQ_554VEOOO_WNEJYCHJ_KE5YKKZN_KFVQX22D_5YHV2HXM",
			AccountBalance: 0,
			NodePublicKey: []byte{55, 227, 166, 111, 37, 197, 148, 220, 78, 154, 194, 204, 113, 162, 167, 78, 167, 246, 146,
				183, 211, 149, 252, 126, 169, 74, 122, 135, 160, 120, 132, 116},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_AHMLIEUZ_G5COQOCF_S2YL2K37_4KMKWZLY_DPVALUZ4_WFKXT54Z_LYBAI323",
			AccountBalance: 0,
			NodePublicKey: []byte{118, 87, 83, 239, 252, 82, 58, 79, 10, 239, 204, 224, 92, 31, 127, 140, 240, 152, 218,
				188, 191, 241, 243, 248, 233, 37, 122, 211, 119, 79, 194, 224},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_XEYICOW4_PM4TCV2K_PFZWMMAM_F5HBAI6E_NM5BQ7CL_WMOZPLIS_ZXNA4PE4",
			AccountBalance: 0,
			NodePublicKey: []byte{48, 45, 138, 55, 36, 219, 105, 79, 139, 39, 225, 21, 236, 200, 10, 237, 92, 236, 28,
				137, 34, 188, 104, 47, 46, 75, 200, 105, 15, 103, 9, 71},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_VXI3BVVO_UFSOSWOD_LFR6L2UX_3OXLIM7R_VJOATCEA_VLPQSJRV_X4TOZO3I",
			AccountBalance: 0,
			NodePublicKey: []byte{103, 58, 238, 203, 109, 188, 122, 99, 180, 168, 139, 141, 75, 187, 181, 27, 174, 194, 16,
				102, 89, 89, 212, 217, 26, 183, 251, 244, 185, 41, 34, 117},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_SZNADILZ_VVRWYHFG_Q24JEDFK_AHA4K6UW_XST46PUX_6VA67MJH_4DKCQRE7",
			AccountBalance: 0,
			NodePublicKey: []byte{127, 76, 216, 115, 18, 65, 180, 92, 57, 218, 105, 1, 50, 11, 121, 143, 243, 188, 195,
				59, 4, 27, 106, 196, 77, 52, 112, 187, 82, 202, 67, 58},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_VXZ4YKM3_UP3PQKLJ_XWEW2ROJ_D3QJ2YLL_YWJQO46J_VKZXGFWS_DQ4JK5TN",
			AccountBalance: 0,
			NodePublicKey: []byte{167, 106, 92, 137, 140, 239, 240, 24, 189, 205, 48, 184, 195, 198, 138, 252, 146, 34, 153,
				65, 56, 52, 133, 104, 179, 200, 82, 50, 142, 154, 128, 60},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_64564FEK_R3CTESGK_PABN72IP_SN47PKGI_P64YTUFZ_ITQFG55N_QBFWSJ7C",
			AccountBalance: 0,
			NodePublicKey: []byte{249, 130, 41, 122, 36, 129, 202, 209, 137, 243, 183, 46, 110, 68, 138, 8, 92, 87, 210,
				195, 12, 202, 171, 179, 77, 131, 42, 111, 180, 148, 108, 144},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_X4Q2AZKH_4GZVJTKX_NOUMMPPN_4XU3APAE_VZINWBUY_3N3PD5RB_TUJH2DT2",
			AccountBalance: 0,
			NodePublicKey: []byte{123, 155, 35, 69, 94, 216, 64, 221, 148, 37, 189, 171, 175, 103, 208, 154, 220, 115, 68,
				4, 64, 115, 43, 167, 243, 176, 98, 66, 75, 45, 200, 106},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_RFNVGNCF_T4R73NSD_YXWMCVID_UJXVE5A4_VAB3MWJK_7N3COA2M_P4VEMAFZ",
			AccountBalance: 0,
			NodePublicKey: []byte{97, 26, 73, 57, 181, 41, 1, 231, 30, 43, 255, 247, 158, 198, 79, 64, 131, 125, 100,
				235, 145, 234, 98, 84, 89, 30, 159, 88, 114, 242, 220, 39},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_66VWZZBG_LPZ4OTWY_ONIVGJDS_FVSSKLX5_LBRUV7LR_GPS4IHVM_AHTSD3GZ",
			AccountBalance: 0,
			NodePublicKey: []byte{29, 243, 69, 231, 243, 13, 51, 239, 228, 10, 29, 130, 26, 107, 223, 19, 72, 73, 64,
				16, 59, 119, 234, 195, 38, 163, 45, 212, 175, 90, 150, 5},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_5KJ7HD4Q_S6WBH3UY_4QZOQTKE_4SHBF6L2_AHHHUTYC_NZFPINZZ_DYXT5BVP",
			AccountBalance: 0,
			NodePublicKey: []byte{148, 17, 123, 166, 88, 172, 192, 15, 101, 184, 242, 105, 228, 129, 99, 223, 228, 44, 203,
				173, 11, 170, 197, 93, 93, 223, 35, 18, 6, 74, 67, 143},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_VBUJJP3J_Y6NIG7SC_OEY2DYR6_RDNS6JYO_SHZ7YFLI_3P5HHKYY_GCPPKXMZ",
			AccountBalance: 0,
			NodePublicKey: []byte{50, 1, 19, 226, 94, 154, 1, 148, 210, 145, 158, 213, 224, 51, 69, 138, 222, 79, 55,
				153, 204, 104, 49, 20, 244, 177, 253, 16, 35, 14, 111, 236},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_7GKXGNBM_MVHDR627_MVKHMGNV_ZDLDFAEA_3KIJ5PGZ_AMC6B4JA_DKNP6HRO",
			AccountBalance: 0,
			NodePublicKey: []byte{243, 150, 36, 123, 156, 102, 181, 163, 152, 245, 22, 8, 21, 120, 132, 141, 42, 194, 117,
				103, 56, 61, 121, 142, 246, 151, 119, 170, 175, 223, 212, 205},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_JW3ZLB7L_ELEH46VI_G7MHAKAB_QHE5ZGZF_J5HVQLHC_OS3NMRGF_XLFJO74K",
			AccountBalance: 0,
			NodePublicKey: []byte{175, 206, 167, 115, 153, 54, 86, 77, 198, 1, 92, 45, 205, 107, 115, 157, 79, 155, 30,
				14, 253, 55, 95, 126, 96, 155, 99, 43, 140, 107, 243, 133},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_AS6KIUHM_B6NAMYUF_23WKK3V7_6VMDPSUF_Y5W3CWOO_SDYAFTZN_HJOCJNVJ",
			AccountBalance: 0,
			NodePublicKey: []byte{136, 234, 45, 1, 94, 205, 145, 136, 87, 131, 220, 113, 176, 99, 18, 150, 144, 97, 145,
				11, 51, 62, 171, 35, 139, 23, 194, 68, 184, 81, 23, 246},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_D4CTM3AT_4EE3RP55_47UIQWNC_NRSJ37L3_3X25ETX4_OVR3KOD5_SSEWNZ6B",
			AccountBalance: 0,
			NodePublicKey: []byte{219, 214, 83, 209, 137, 245, 78, 201, 52, 0, 166, 40, 217, 7, 43, 212, 192, 239, 25,
				42, 245, 220, 155, 225, 26, 105, 174, 142, 121, 24, 176, 124},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_ZDROGLWB_3EBTFWXO_PED4N3DV_H47ZHQWZ_GKRYRIAH_RB5KIGLA_WEDQVHR4",
			AccountBalance: 0,
			NodePublicKey: []byte{239, 189, 154, 171, 216, 19, 3, 125, 107, 49, 166, 134, 181, 75, 48, 55, 97, 87, 143,
				14, 233, 186, 15, 129, 90, 166, 15, 102, 129, 149, 100, 248},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_SBTR4VON_FRNJJHU5_ZXK66XXN_D2MHAXVO_YMPVNFOH_WBD7RHG6_GI2DCEAI",
			AccountBalance: 0,
			NodePublicKey: []byte{219, 252, 85, 53, 170, 52, 228, 142, 26, 192, 41, 102, 97, 123, 18, 158, 180, 150, 231,
				44, 180, 160, 94, 169, 235, 120, 77, 148, 155, 102, 158, 40},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_QPHVOLVS_OMQD73A6_D4TI64FB_BR3QDIXS_XAXD7QAT_UCH3RL3Y_2VPLA6CC",
			AccountBalance: 0,
			NodePublicKey: []byte{145, 213, 162, 234, 160, 67, 216, 89, 157, 115, 79, 255, 105, 43, 183, 205, 15, 19, 66,
				210, 54, 60, 208, 151, 90, 219, 133, 197, 217, 178, 39, 149},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_3CVD7YDM_56GETU2R_IIZ7ZFN6_X7HP5G2C_ZZCHYZIL_QURRMVTJ_C5DCZIR4",
			AccountBalance: 0,
			NodePublicKey: []byte{237, 204, 215, 142, 252, 85, 7, 13, 61, 127, 81, 233, 106, 254, 9, 237, 251, 210, 48,
				53, 132, 176, 156, 118, 102, 47, 44, 1, 236, 108, 146, 217},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_YCE35IPA_4KIZMSQF_DQPECAKA_GMXC33SH_4ILSHEZT_GL3WCRLO_UX6CGJP2",
			AccountBalance: 0,
			NodePublicKey: []byte{151, 249, 110, 3, 81, 35, 247, 47, 129, 13, 251, 243, 254, 49, 214, 230, 86, 250, 106,
				219, 247, 114, 145, 175, 175, 246, 124, 169, 44, 131, 229, 158},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_K52C2E5W_5ODKDJQ6_WNNW7F36_OMIOO6QH_N52OCQ3E_Z4VVJMYF_5WUUPFAG",
			AccountBalance: 0,
			NodePublicKey: []byte{40, 26, 59, 198, 120, 194, 87, 140, 23, 58, 14, 48, 215, 34, 70, 121, 154, 50, 162,
				128, 186, 248, 87, 187, 153, 54, 137, 143, 18, 147, 146, 138},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_FZMA7JWY_B6KFGZUJ_HUVPEPD3_SNUNIBJQ_WHBELID3_VONMUOS4_EWBBRA22",
			AccountBalance: 0,
			NodePublicKey: []byte{188, 94, 192, 45, 30, 100, 206, 179, 239, 230, 93, 195, 137, 68, 165, 69, 142, 186, 133,
				209, 125, 83, 128, 230, 208, 207, 82, 48, 51, 147, 69, 122},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
		{
			AccountAddress: "ZBC_YZ4OHUXU_BHRIF4S2_TOZWWBGM_AW7ETJJ5_OBBZD5UT_ND7BNVHA_FWHEAYTQ",
			AccountBalance: 0,
			NodePublicKey: []byte{230, 58, 65, 142, 239, 162, 249, 122, 107, 242, 184, 117, 111, 185, 21, 17, 101, 224, 178,
				173, 188, 155, 210, 167, 67, 142, 105, 183, 6, 22, 163, 62},
			LockedBalance:      0,
			ParticipationScore: GenesisParticipationScore,
		},
	}
)
