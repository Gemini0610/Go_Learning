// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"demo/internal/dao/internal"
)

// internalCouponInfoDao is internal type for wrapping internal DAO implements.
type internalCouponInfoDao = *internal.CouponInfoDao

// couponInfoDao is the data access object for table coupon_info.
// You can define custom methods on it to extend its functionality as you wish.
type couponInfoDao struct {
	internalCouponInfoDao
}

var (
	// CouponInfo is globally public accessible object for table coupon_info operations.
	CouponInfo = couponInfoDao{
		internal.NewCouponInfoDao(),
	}
)

// Fill with you ideas below.
