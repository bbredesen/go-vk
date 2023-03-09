// Code generated by "stringer -output=enum_string_1.go -type=ComponentSwizzle,ComponentTypeNV,CompositeAlphaFlagBitsKHR,ConditionalRenderingFlagBitsEXT,ConservativeRasterizationModeEXT,CopyAccelerationStructureModeKHR,CopyMicromapModeEXT,CoverageModulationModeNV,CoverageReductionModeNV,CullModeFlagBits,DebugReportFlagBitsEXT,DebugReportObjectTypeEXT,DebugUtilsMessageSeverityFlagBitsEXT,DebugUtilsMessageTypeFlagBitsEXT,DependencyFlagBits,DescriptorBindingFlagBits,DescriptorPoolCreateFlagBits,DescriptorSetLayoutCreateFlagBits,DescriptorType,DescriptorUpdateTemplateType,DeviceAddressBindingFlagBitsEXT,DeviceAddressBindingTypeEXT,DeviceDiagnosticsConfigFlagBitsNV,DeviceEventTypeEXT,DeviceFaultAddressTypeEXT,DeviceFaultVendorBinaryHeaderVersionEXT,DeviceGroupPresentModeFlagBitsKHR,DeviceMemoryReportEventTypeEXT,DeviceQueueCreateFlagBits,DirectDriverLoadingModeLUNARG,DiscardRectangleModeEXT"; DO NOT EDIT.

package vk

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[COMPONENT_SWIZZLE_IDENTITY-0]
	_ = x[COMPONENT_SWIZZLE_ZERO-1]
	_ = x[COMPONENT_SWIZZLE_ONE-2]
	_ = x[COMPONENT_SWIZZLE_R-3]
	_ = x[COMPONENT_SWIZZLE_G-4]
	_ = x[COMPONENT_SWIZZLE_B-5]
	_ = x[COMPONENT_SWIZZLE_A-6]
}

const _ComponentSwizzle_name = "COMPONENT_SWIZZLE_IDENTITYCOMPONENT_SWIZZLE_ZEROCOMPONENT_SWIZZLE_ONECOMPONENT_SWIZZLE_RCOMPONENT_SWIZZLE_GCOMPONENT_SWIZZLE_BCOMPONENT_SWIZZLE_A"

var _ComponentSwizzle_index = [...]uint8{0, 26, 48, 69, 88, 107, 126, 145}

func (i ComponentSwizzle) String() string {
	if i < 0 || i >= ComponentSwizzle(len(_ComponentSwizzle_index)-1) {
		return "ComponentSwizzle(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ComponentSwizzle_name[_ComponentSwizzle_index[i]:_ComponentSwizzle_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[COMPONENT_TYPE_FLOAT16_NV-0]
	_ = x[COMPONENT_TYPE_FLOAT32_NV-1]
	_ = x[COMPONENT_TYPE_FLOAT64_NV-2]
	_ = x[COMPONENT_TYPE_SINT8_NV-3]
	_ = x[COMPONENT_TYPE_SINT16_NV-4]
	_ = x[COMPONENT_TYPE_SINT32_NV-5]
	_ = x[COMPONENT_TYPE_SINT64_NV-6]
	_ = x[COMPONENT_TYPE_UINT8_NV-7]
	_ = x[COMPONENT_TYPE_UINT16_NV-8]
	_ = x[COMPONENT_TYPE_UINT32_NV-9]
	_ = x[COMPONENT_TYPE_UINT64_NV-10]
}

const _ComponentTypeNV_name = "COMPONENT_TYPE_FLOAT16_NVCOMPONENT_TYPE_FLOAT32_NVCOMPONENT_TYPE_FLOAT64_NVCOMPONENT_TYPE_SINT8_NVCOMPONENT_TYPE_SINT16_NVCOMPONENT_TYPE_SINT32_NVCOMPONENT_TYPE_SINT64_NVCOMPONENT_TYPE_UINT8_NVCOMPONENT_TYPE_UINT16_NVCOMPONENT_TYPE_UINT32_NVCOMPONENT_TYPE_UINT64_NV"

var _ComponentTypeNV_index = [...]uint16{0, 25, 50, 75, 98, 122, 146, 170, 193, 217, 241, 265}

func (i ComponentTypeNV) String() string {
	if i < 0 || i >= ComponentTypeNV(len(_ComponentTypeNV_index)-1) {
		return "ComponentTypeNV(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ComponentTypeNV_name[_ComponentTypeNV_index[i]:_ComponentTypeNV_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[COMPOSITE_ALPHA_OPAQUE_BIT_KHR-1]
	_ = x[COMPOSITE_ALPHA_PRE_MULTIPLIED_BIT_KHR-2]
	_ = x[COMPOSITE_ALPHA_POST_MULTIPLIED_BIT_KHR-4]
	_ = x[COMPOSITE_ALPHA_INHERIT_BIT_KHR-8]
}

const (
	_CompositeAlphaFlagBitsKHR_name_0 = "COMPOSITE_ALPHA_OPAQUE_BIT_KHRCOMPOSITE_ALPHA_PRE_MULTIPLIED_BIT_KHR"
	_CompositeAlphaFlagBitsKHR_name_1 = "COMPOSITE_ALPHA_POST_MULTIPLIED_BIT_KHR"
	_CompositeAlphaFlagBitsKHR_name_2 = "COMPOSITE_ALPHA_INHERIT_BIT_KHR"
)

var (
	_CompositeAlphaFlagBitsKHR_index_0 = [...]uint8{0, 30, 68}
)

func (i CompositeAlphaFlagBitsKHR) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _CompositeAlphaFlagBitsKHR_name_0[_CompositeAlphaFlagBitsKHR_index_0[i]:_CompositeAlphaFlagBitsKHR_index_0[i+1]]
	case i == 4:
		return _CompositeAlphaFlagBitsKHR_name_1
	case i == 8:
		return _CompositeAlphaFlagBitsKHR_name_2
	default:
		return "CompositeAlphaFlagBitsKHR(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CONDITIONAL_RENDERING_INVERTED_BIT_EXT-1]
}

const _ConditionalRenderingFlagBitsEXT_name = "CONDITIONAL_RENDERING_INVERTED_BIT_EXT"

var _ConditionalRenderingFlagBitsEXT_index = [...]uint8{0, 38}

func (i ConditionalRenderingFlagBitsEXT) String() string {
	i -= 1
	if i >= ConditionalRenderingFlagBitsEXT(len(_ConditionalRenderingFlagBitsEXT_index)-1) {
		return "ConditionalRenderingFlagBitsEXT(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _ConditionalRenderingFlagBitsEXT_name[_ConditionalRenderingFlagBitsEXT_index[i]:_ConditionalRenderingFlagBitsEXT_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CONSERVATIVE_RASTERIZATION_MODE_DISABLED_EXT-0]
	_ = x[CONSERVATIVE_RASTERIZATION_MODE_OVERESTIMATE_EXT-1]
	_ = x[CONSERVATIVE_RASTERIZATION_MODE_UNDERESTIMATE_EXT-2]
}

const _ConservativeRasterizationModeEXT_name = "CONSERVATIVE_RASTERIZATION_MODE_DISABLED_EXTCONSERVATIVE_RASTERIZATION_MODE_OVERESTIMATE_EXTCONSERVATIVE_RASTERIZATION_MODE_UNDERESTIMATE_EXT"

var _ConservativeRasterizationModeEXT_index = [...]uint8{0, 44, 92, 141}

func (i ConservativeRasterizationModeEXT) String() string {
	if i < 0 || i >= ConservativeRasterizationModeEXT(len(_ConservativeRasterizationModeEXT_index)-1) {
		return "ConservativeRasterizationModeEXT(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _ConservativeRasterizationModeEXT_name[_ConservativeRasterizationModeEXT_index[i]:_ConservativeRasterizationModeEXT_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[COPY_ACCELERATION_STRUCTURE_MODE_CLONE_KHR-0]
	_ = x[COPY_ACCELERATION_STRUCTURE_MODE_COMPACT_KHR-1]
	_ = x[COPY_ACCELERATION_STRUCTURE_MODE_SERIALIZE_KHR-2]
	_ = x[COPY_ACCELERATION_STRUCTURE_MODE_DESERIALIZE_KHR-3]
	_ = x[COPY_ACCELERATION_STRUCTURE_MODE_CLONE_NV-0]
	_ = x[COPY_ACCELERATION_STRUCTURE_MODE_COMPACT_NV-1]
}

const _CopyAccelerationStructureModeKHR_name = "COPY_ACCELERATION_STRUCTURE_MODE_CLONE_KHRCOPY_ACCELERATION_STRUCTURE_MODE_COMPACT_KHRCOPY_ACCELERATION_STRUCTURE_MODE_SERIALIZE_KHRCOPY_ACCELERATION_STRUCTURE_MODE_DESERIALIZE_KHR"

var _CopyAccelerationStructureModeKHR_index = [...]uint8{0, 42, 86, 132, 180}

func (i CopyAccelerationStructureModeKHR) String() string {
	if i < 0 || i >= CopyAccelerationStructureModeKHR(len(_CopyAccelerationStructureModeKHR_index)-1) {
		return "CopyAccelerationStructureModeKHR(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CopyAccelerationStructureModeKHR_name[_CopyAccelerationStructureModeKHR_index[i]:_CopyAccelerationStructureModeKHR_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[COPY_MICROMAP_MODE_CLONE_EXT-0]
	_ = x[COPY_MICROMAP_MODE_SERIALIZE_EXT-1]
	_ = x[COPY_MICROMAP_MODE_DESERIALIZE_EXT-2]
	_ = x[COPY_MICROMAP_MODE_COMPACT_EXT-3]
}

const _CopyMicromapModeEXT_name = "COPY_MICROMAP_MODE_CLONE_EXTCOPY_MICROMAP_MODE_SERIALIZE_EXTCOPY_MICROMAP_MODE_DESERIALIZE_EXTCOPY_MICROMAP_MODE_COMPACT_EXT"

var _CopyMicromapModeEXT_index = [...]uint8{0, 28, 60, 94, 124}

func (i CopyMicromapModeEXT) String() string {
	if i < 0 || i >= CopyMicromapModeEXT(len(_CopyMicromapModeEXT_index)-1) {
		return "CopyMicromapModeEXT(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CopyMicromapModeEXT_name[_CopyMicromapModeEXT_index[i]:_CopyMicromapModeEXT_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[COVERAGE_MODULATION_MODE_NONE_NV-0]
	_ = x[COVERAGE_MODULATION_MODE_RGB_NV-1]
	_ = x[COVERAGE_MODULATION_MODE_ALPHA_NV-2]
	_ = x[COVERAGE_MODULATION_MODE_RGBA_NV-3]
}

const _CoverageModulationModeNV_name = "COVERAGE_MODULATION_MODE_NONE_NVCOVERAGE_MODULATION_MODE_RGB_NVCOVERAGE_MODULATION_MODE_ALPHA_NVCOVERAGE_MODULATION_MODE_RGBA_NV"

var _CoverageModulationModeNV_index = [...]uint8{0, 32, 63, 96, 128}

func (i CoverageModulationModeNV) String() string {
	if i < 0 || i >= CoverageModulationModeNV(len(_CoverageModulationModeNV_index)-1) {
		return "CoverageModulationModeNV(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CoverageModulationModeNV_name[_CoverageModulationModeNV_index[i]:_CoverageModulationModeNV_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[COVERAGE_REDUCTION_MODE_MERGE_NV-0]
	_ = x[COVERAGE_REDUCTION_MODE_TRUNCATE_NV-1]
}

const _CoverageReductionModeNV_name = "COVERAGE_REDUCTION_MODE_MERGE_NVCOVERAGE_REDUCTION_MODE_TRUNCATE_NV"

var _CoverageReductionModeNV_index = [...]uint8{0, 32, 67}

func (i CoverageReductionModeNV) String() string {
	if i < 0 || i >= CoverageReductionModeNV(len(_CoverageReductionModeNV_index)-1) {
		return "CoverageReductionModeNV(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CoverageReductionModeNV_name[_CoverageReductionModeNV_index[i]:_CoverageReductionModeNV_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[CULL_MODE_NONE-0]
	_ = x[CULL_MODE_FRONT_AND_BACK-3]
	_ = x[CULL_MODE_FRONT_BIT-1]
	_ = x[CULL_MODE_BACK_BIT-2]
}

const _CullModeFlagBits_name = "CULL_MODE_NONECULL_MODE_FRONT_BITCULL_MODE_BACK_BITCULL_MODE_FRONT_AND_BACK"

var _CullModeFlagBits_index = [...]uint8{0, 14, 33, 51, 75}

func (i CullModeFlagBits) String() string {
	if i >= CullModeFlagBits(len(_CullModeFlagBits_index)-1) {
		return "CullModeFlagBits(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _CullModeFlagBits_name[_CullModeFlagBits_index[i]:_CullModeFlagBits_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DEBUG_REPORT_INFORMATION_BIT_EXT-1]
	_ = x[DEBUG_REPORT_WARNING_BIT_EXT-2]
	_ = x[DEBUG_REPORT_PERFORMANCE_WARNING_BIT_EXT-4]
	_ = x[DEBUG_REPORT_ERROR_BIT_EXT-8]
	_ = x[DEBUG_REPORT_DEBUG_BIT_EXT-16]
}

const (
	_DebugReportFlagBitsEXT_name_0 = "DEBUG_REPORT_INFORMATION_BIT_EXTDEBUG_REPORT_WARNING_BIT_EXT"
	_DebugReportFlagBitsEXT_name_1 = "DEBUG_REPORT_PERFORMANCE_WARNING_BIT_EXT"
	_DebugReportFlagBitsEXT_name_2 = "DEBUG_REPORT_ERROR_BIT_EXT"
	_DebugReportFlagBitsEXT_name_3 = "DEBUG_REPORT_DEBUG_BIT_EXT"
)

var (
	_DebugReportFlagBitsEXT_index_0 = [...]uint8{0, 32, 60}
)

func (i DebugReportFlagBitsEXT) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _DebugReportFlagBitsEXT_name_0[_DebugReportFlagBitsEXT_index_0[i]:_DebugReportFlagBitsEXT_index_0[i+1]]
	case i == 4:
		return _DebugReportFlagBitsEXT_name_1
	case i == 8:
		return _DebugReportFlagBitsEXT_name_2
	case i == 16:
		return _DebugReportFlagBitsEXT_name_3
	default:
		return "DebugReportFlagBitsEXT(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DEBUG_REPORT_OBJECT_TYPE_UNKNOWN_EXT-0]
	_ = x[DEBUG_REPORT_OBJECT_TYPE_INSTANCE_EXT-1]
	_ = x[DEBUG_REPORT_OBJECT_TYPE_PHYSICAL_DEVICE_EXT-2]
	_ = x[DEBUG_REPORT_OBJECT_TYPE_DEVICE_EXT-3]
	_ = x[DEBUG_REPORT_OBJECT_TYPE_QUEUE_EXT-4]
	_ = x[DEBUG_REPORT_OBJECT_TYPE_SEMAPHORE_EXT-5]
	_ = x[DEBUG_REPORT_OBJECT_TYPE_COMMAND_BUFFER_EXT-6]
	_ = x[DEBUG_REPORT_OBJECT_TYPE_FENCE_EXT-7]
	_ = x[DEBUG_REPORT_OBJECT_TYPE_DEVICE_MEMORY_EXT-8]
	_ = x[DEBUG_REPORT_OBJECT_TYPE_BUFFER_EXT-9]
	_ = x[DEBUG_REPORT_OBJECT_TYPE_IMAGE_EXT-10]
	_ = x[DEBUG_REPORT_OBJECT_TYPE_EVENT_EXT-11]
	_ = x[DEBUG_REPORT_OBJECT_TYPE_QUERY_POOL_EXT-12]
	_ = x[DEBUG_REPORT_OBJECT_TYPE_BUFFER_VIEW_EXT-13]
	_ = x[DEBUG_REPORT_OBJECT_TYPE_IMAGE_VIEW_EXT-14]
	_ = x[DEBUG_REPORT_OBJECT_TYPE_SHADER_MODULE_EXT-15]
	_ = x[DEBUG_REPORT_OBJECT_TYPE_PIPELINE_CACHE_EXT-16]
	_ = x[DEBUG_REPORT_OBJECT_TYPE_PIPELINE_LAYOUT_EXT-17]
	_ = x[DEBUG_REPORT_OBJECT_TYPE_RENDER_PASS_EXT-18]
	_ = x[DEBUG_REPORT_OBJECT_TYPE_PIPELINE_EXT-19]
	_ = x[DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_SET_LAYOUT_EXT-20]
	_ = x[DEBUG_REPORT_OBJECT_TYPE_SAMPLER_EXT-21]
	_ = x[DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_POOL_EXT-22]
	_ = x[DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_SET_EXT-23]
	_ = x[DEBUG_REPORT_OBJECT_TYPE_FRAMEBUFFER_EXT-24]
	_ = x[DEBUG_REPORT_OBJECT_TYPE_COMMAND_POOL_EXT-25]
	_ = x[DEBUG_REPORT_OBJECT_TYPE_SURFACE_KHR_EXT-26]
	_ = x[DEBUG_REPORT_OBJECT_TYPE_SWAPCHAIN_KHR_EXT-27]
	_ = x[DEBUG_REPORT_OBJECT_TYPE_DEBUG_REPORT_CALLBACK_EXT_EXT-28]
	_ = x[DEBUG_REPORT_OBJECT_TYPE_DISPLAY_KHR_EXT-29]
	_ = x[DEBUG_REPORT_OBJECT_TYPE_DISPLAY_MODE_KHR_EXT-30]
	_ = x[DEBUG_REPORT_OBJECT_TYPE_VALIDATION_CACHE_EXT_EXT-33]
	_ = x[DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_EXT-1000011000]
	_ = x[DEBUG_REPORT_OBJECT_TYPE_CU_MODULE_NVX_EXT-1000029000]
	_ = x[DEBUG_REPORT_OBJECT_TYPE_CU_FUNCTION_NVX_EXT-1000029001]
	_ = x[DEBUG_REPORT_OBJECT_TYPE_ACCELERATION_STRUCTURE_KHR_EXT-1000150000]
	_ = x[DEBUG_REPORT_OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION_EXT-1000156000]
	_ = x[DEBUG_REPORT_OBJECT_TYPE_ACCELERATION_STRUCTURE_NV_EXT-1000165000]
	_ = x[DEBUG_REPORT_OBJECT_TYPE_DEBUG_REPORT_EXT-28]
	_ = x[DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_KHR_EXT-1000011000]
	_ = x[DEBUG_REPORT_OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION_KHR_EXT-1000156000]
	_ = x[DEBUG_REPORT_OBJECT_TYPE_VALIDATION_CACHE_EXT-33]
}

const (
	_DebugReportObjectTypeEXT_name_0 = "DEBUG_REPORT_OBJECT_TYPE_UNKNOWN_EXTDEBUG_REPORT_OBJECT_TYPE_INSTANCE_EXTDEBUG_REPORT_OBJECT_TYPE_PHYSICAL_DEVICE_EXTDEBUG_REPORT_OBJECT_TYPE_DEVICE_EXTDEBUG_REPORT_OBJECT_TYPE_QUEUE_EXTDEBUG_REPORT_OBJECT_TYPE_SEMAPHORE_EXTDEBUG_REPORT_OBJECT_TYPE_COMMAND_BUFFER_EXTDEBUG_REPORT_OBJECT_TYPE_FENCE_EXTDEBUG_REPORT_OBJECT_TYPE_DEVICE_MEMORY_EXTDEBUG_REPORT_OBJECT_TYPE_BUFFER_EXTDEBUG_REPORT_OBJECT_TYPE_IMAGE_EXTDEBUG_REPORT_OBJECT_TYPE_EVENT_EXTDEBUG_REPORT_OBJECT_TYPE_QUERY_POOL_EXTDEBUG_REPORT_OBJECT_TYPE_BUFFER_VIEW_EXTDEBUG_REPORT_OBJECT_TYPE_IMAGE_VIEW_EXTDEBUG_REPORT_OBJECT_TYPE_SHADER_MODULE_EXTDEBUG_REPORT_OBJECT_TYPE_PIPELINE_CACHE_EXTDEBUG_REPORT_OBJECT_TYPE_PIPELINE_LAYOUT_EXTDEBUG_REPORT_OBJECT_TYPE_RENDER_PASS_EXTDEBUG_REPORT_OBJECT_TYPE_PIPELINE_EXTDEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_SET_LAYOUT_EXTDEBUG_REPORT_OBJECT_TYPE_SAMPLER_EXTDEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_POOL_EXTDEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_SET_EXTDEBUG_REPORT_OBJECT_TYPE_FRAMEBUFFER_EXTDEBUG_REPORT_OBJECT_TYPE_COMMAND_POOL_EXTDEBUG_REPORT_OBJECT_TYPE_SURFACE_KHR_EXTDEBUG_REPORT_OBJECT_TYPE_SWAPCHAIN_KHR_EXTDEBUG_REPORT_OBJECT_TYPE_DEBUG_REPORT_CALLBACK_EXT_EXTDEBUG_REPORT_OBJECT_TYPE_DISPLAY_KHR_EXTDEBUG_REPORT_OBJECT_TYPE_DISPLAY_MODE_KHR_EXT"
	_DebugReportObjectTypeEXT_name_1 = "DEBUG_REPORT_OBJECT_TYPE_VALIDATION_CACHE_EXT_EXT"
	_DebugReportObjectTypeEXT_name_2 = "DEBUG_REPORT_OBJECT_TYPE_DESCRIPTOR_UPDATE_TEMPLATE_EXT"
	_DebugReportObjectTypeEXT_name_3 = "DEBUG_REPORT_OBJECT_TYPE_CU_MODULE_NVX_EXTDEBUG_REPORT_OBJECT_TYPE_CU_FUNCTION_NVX_EXT"
	_DebugReportObjectTypeEXT_name_4 = "DEBUG_REPORT_OBJECT_TYPE_ACCELERATION_STRUCTURE_KHR_EXT"
	_DebugReportObjectTypeEXT_name_5 = "DEBUG_REPORT_OBJECT_TYPE_SAMPLER_YCBCR_CONVERSION_EXT"
	_DebugReportObjectTypeEXT_name_6 = "DEBUG_REPORT_OBJECT_TYPE_ACCELERATION_STRUCTURE_NV_EXT"
)

var (
	_DebugReportObjectTypeEXT_index_0 = [...]uint16{0, 36, 73, 117, 152, 186, 224, 267, 301, 343, 378, 412, 446, 485, 525, 564, 606, 649, 693, 733, 770, 820, 856, 900, 943, 983, 1024, 1064, 1106, 1160, 1200, 1245}
	_DebugReportObjectTypeEXT_index_3 = [...]uint8{0, 42, 86}
)

func (i DebugReportObjectTypeEXT) String() string {
	switch {
	case 0 <= i && i <= 30:
		return _DebugReportObjectTypeEXT_name_0[_DebugReportObjectTypeEXT_index_0[i]:_DebugReportObjectTypeEXT_index_0[i+1]]
	case i == 33:
		return _DebugReportObjectTypeEXT_name_1
	case i == 1000011000:
		return _DebugReportObjectTypeEXT_name_2
	case 1000029000 <= i && i <= 1000029001:
		i -= 1000029000
		return _DebugReportObjectTypeEXT_name_3[_DebugReportObjectTypeEXT_index_3[i]:_DebugReportObjectTypeEXT_index_3[i+1]]
	case i == 1000150000:
		return _DebugReportObjectTypeEXT_name_4
	case i == 1000156000:
		return _DebugReportObjectTypeEXT_name_5
	case i == 1000165000:
		return _DebugReportObjectTypeEXT_name_6
	default:
		return "DebugReportObjectTypeEXT(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DEBUG_UTILS_MESSAGE_SEVERITY_VERBOSE_BIT_EXT-1]
	_ = x[DEBUG_UTILS_MESSAGE_SEVERITY_ERROR_BIT_EXT-4096]
	_ = x[DEBUG_UTILS_MESSAGE_SEVERITY_INFO_BIT_EXT-16]
	_ = x[DEBUG_UTILS_MESSAGE_SEVERITY_WARNING_BIT_EXT-256]
}

const (
	_DebugUtilsMessageSeverityFlagBitsEXT_name_0 = "DEBUG_UTILS_MESSAGE_SEVERITY_VERBOSE_BIT_EXT"
	_DebugUtilsMessageSeverityFlagBitsEXT_name_1 = "DEBUG_UTILS_MESSAGE_SEVERITY_INFO_BIT_EXT"
	_DebugUtilsMessageSeverityFlagBitsEXT_name_2 = "DEBUG_UTILS_MESSAGE_SEVERITY_WARNING_BIT_EXT"
	_DebugUtilsMessageSeverityFlagBitsEXT_name_3 = "DEBUG_UTILS_MESSAGE_SEVERITY_ERROR_BIT_EXT"
)

func (i DebugUtilsMessageSeverityFlagBitsEXT) String() string {
	switch {
	case i == 1:
		return _DebugUtilsMessageSeverityFlagBitsEXT_name_0
	case i == 16:
		return _DebugUtilsMessageSeverityFlagBitsEXT_name_1
	case i == 256:
		return _DebugUtilsMessageSeverityFlagBitsEXT_name_2
	case i == 4096:
		return _DebugUtilsMessageSeverityFlagBitsEXT_name_3
	default:
		return "DebugUtilsMessageSeverityFlagBitsEXT(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DEBUG_UTILS_MESSAGE_TYPE_GENERAL_BIT_EXT-1]
	_ = x[DEBUG_UTILS_MESSAGE_TYPE_VALIDATION_BIT_EXT-2]
	_ = x[DEBUG_UTILS_MESSAGE_TYPE_PERFORMANCE_BIT_EXT-4]
	_ = x[DEBUG_UTILS_MESSAGE_TYPE_DEVICE_ADDRESS_BINDING_BIT_EXT-1000354000]
}

const (
	_DebugUtilsMessageTypeFlagBitsEXT_name_0 = "DEBUG_UTILS_MESSAGE_TYPE_GENERAL_BIT_EXTDEBUG_UTILS_MESSAGE_TYPE_VALIDATION_BIT_EXT"
	_DebugUtilsMessageTypeFlagBitsEXT_name_1 = "DEBUG_UTILS_MESSAGE_TYPE_PERFORMANCE_BIT_EXT"
	_DebugUtilsMessageTypeFlagBitsEXT_name_2 = "DEBUG_UTILS_MESSAGE_TYPE_DEVICE_ADDRESS_BINDING_BIT_EXT"
)

var (
	_DebugUtilsMessageTypeFlagBitsEXT_index_0 = [...]uint8{0, 40, 83}
)

func (i DebugUtilsMessageTypeFlagBitsEXT) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _DebugUtilsMessageTypeFlagBitsEXT_name_0[_DebugUtilsMessageTypeFlagBitsEXT_index_0[i]:_DebugUtilsMessageTypeFlagBitsEXT_index_0[i+1]]
	case i == 4:
		return _DebugUtilsMessageTypeFlagBitsEXT_name_1
	case i == 1000354000:
		return _DebugUtilsMessageTypeFlagBitsEXT_name_2
	default:
		return "DebugUtilsMessageTypeFlagBitsEXT(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DEPENDENCY_BY_REGION_BIT-1]
	_ = x[DEPENDENCY_VIEW_LOCAL_BIT-2]
	_ = x[DEPENDENCY_DEVICE_GROUP_BIT-4]
	_ = x[DEPENDENCY_FEEDBACK_LOOP_BIT_EXT-1000339000]
	_ = x[DEPENDENCY_DEVICE_GROUP_BIT_KHR-4]
	_ = x[DEPENDENCY_VIEW_LOCAL_BIT_KHR-2]
}

const (
	_DependencyFlagBits_name_0 = "DEPENDENCY_BY_REGION_BITDEPENDENCY_VIEW_LOCAL_BIT"
	_DependencyFlagBits_name_1 = "DEPENDENCY_DEVICE_GROUP_BIT"
	_DependencyFlagBits_name_2 = "DEPENDENCY_FEEDBACK_LOOP_BIT_EXT"
)

var (
	_DependencyFlagBits_index_0 = [...]uint8{0, 24, 49}
)

func (i DependencyFlagBits) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _DependencyFlagBits_name_0[_DependencyFlagBits_index_0[i]:_DependencyFlagBits_index_0[i+1]]
	case i == 4:
		return _DependencyFlagBits_name_1
	case i == 1000339000:
		return _DependencyFlagBits_name_2
	default:
		return "DependencyFlagBits(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT-1]
	_ = x[DESCRIPTOR_BINDING_UPDATE_UNUSED_WHILE_PENDING_BIT-2]
	_ = x[DESCRIPTOR_BINDING_PARTIALLY_BOUND_BIT-4]
	_ = x[DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT-8]
	_ = x[DESCRIPTOR_BINDING_PARTIALLY_BOUND_BIT_EXT-4]
	_ = x[DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BIT_EXT-1]
	_ = x[DESCRIPTOR_BINDING_UPDATE_UNUSED_WHILE_PENDING_BIT_EXT-2]
	_ = x[DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT_EXT-8]
}

const (
	_DescriptorBindingFlagBits_name_0 = "DESCRIPTOR_BINDING_UPDATE_AFTER_BIND_BITDESCRIPTOR_BINDING_UPDATE_UNUSED_WHILE_PENDING_BIT"
	_DescriptorBindingFlagBits_name_1 = "DESCRIPTOR_BINDING_PARTIALLY_BOUND_BIT"
	_DescriptorBindingFlagBits_name_2 = "DESCRIPTOR_BINDING_VARIABLE_DESCRIPTOR_COUNT_BIT"
)

var (
	_DescriptorBindingFlagBits_index_0 = [...]uint8{0, 40, 90}
)

func (i DescriptorBindingFlagBits) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _DescriptorBindingFlagBits_name_0[_DescriptorBindingFlagBits_index_0[i]:_DescriptorBindingFlagBits_index_0[i+1]]
	case i == 4:
		return _DescriptorBindingFlagBits_name_1
	case i == 8:
		return _DescriptorBindingFlagBits_name_2
	default:
		return "DescriptorBindingFlagBits(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DESCRIPTOR_POOL_CREATE_FREE_DESCRIPTOR_SET_BIT-1]
	_ = x[DESCRIPTOR_POOL_CREATE_UPDATE_AFTER_BIND_BIT-2]
	_ = x[DESCRIPTOR_POOL_CREATE_HOST_ONLY_BIT_EXT-1000494000]
	_ = x[DESCRIPTOR_POOL_CREATE_HOST_ONLY_BIT_VALVE-1000494000]
	_ = x[DESCRIPTOR_POOL_CREATE_UPDATE_AFTER_BIND_BIT_EXT-2]
}

const (
	_DescriptorPoolCreateFlagBits_name_0 = "DESCRIPTOR_POOL_CREATE_FREE_DESCRIPTOR_SET_BITDESCRIPTOR_POOL_CREATE_UPDATE_AFTER_BIND_BIT"
	_DescriptorPoolCreateFlagBits_name_1 = "DESCRIPTOR_POOL_CREATE_HOST_ONLY_BIT_EXT"
)

var (
	_DescriptorPoolCreateFlagBits_index_0 = [...]uint8{0, 46, 90}
)

func (i DescriptorPoolCreateFlagBits) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _DescriptorPoolCreateFlagBits_name_0[_DescriptorPoolCreateFlagBits_index_0[i]:_DescriptorPoolCreateFlagBits_index_0[i+1]]
	case i == 1000494000:
		return _DescriptorPoolCreateFlagBits_name_1
	default:
		return "DescriptorPoolCreateFlagBits(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT-2]
	_ = x[DESCRIPTOR_SET_LAYOUT_CREATE_PUSH_DESCRIPTOR_BIT_KHR-1000080000]
	_ = x[DESCRIPTOR_SET_LAYOUT_CREATE_DESCRIPTOR_BUFFER_BIT_EXT-1000316000]
	_ = x[DESCRIPTOR_SET_LAYOUT_CREATE_EMBEDDED_IMMUTABLE_SAMPLERS_BIT_EXT-1000316000]
	_ = x[DESCRIPTOR_SET_LAYOUT_CREATE_HOST_ONLY_POOL_BIT_EXT-1000494000]
	_ = x[DESCRIPTOR_SET_LAYOUT_CREATE_HOST_ONLY_POOL_BIT_VALVE-1000494000]
	_ = x[DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT_EXT-2]
}

const (
	_DescriptorSetLayoutCreateFlagBits_name_0 = "DESCRIPTOR_SET_LAYOUT_CREATE_UPDATE_AFTER_BIND_POOL_BIT"
	_DescriptorSetLayoutCreateFlagBits_name_1 = "DESCRIPTOR_SET_LAYOUT_CREATE_PUSH_DESCRIPTOR_BIT_KHR"
	_DescriptorSetLayoutCreateFlagBits_name_2 = "DESCRIPTOR_SET_LAYOUT_CREATE_DESCRIPTOR_BUFFER_BIT_EXT"
	_DescriptorSetLayoutCreateFlagBits_name_3 = "DESCRIPTOR_SET_LAYOUT_CREATE_HOST_ONLY_POOL_BIT_EXT"
)

func (i DescriptorSetLayoutCreateFlagBits) String() string {
	switch {
	case i == 2:
		return _DescriptorSetLayoutCreateFlagBits_name_0
	case i == 1000080000:
		return _DescriptorSetLayoutCreateFlagBits_name_1
	case i == 1000316000:
		return _DescriptorSetLayoutCreateFlagBits_name_2
	case i == 1000494000:
		return _DescriptorSetLayoutCreateFlagBits_name_3
	default:
		return "DescriptorSetLayoutCreateFlagBits(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DESCRIPTOR_TYPE_SAMPLER-0]
	_ = x[DESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLER-1]
	_ = x[DESCRIPTOR_TYPE_SAMPLED_IMAGE-2]
	_ = x[DESCRIPTOR_TYPE_STORAGE_IMAGE-3]
	_ = x[DESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFER-4]
	_ = x[DESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFER-5]
	_ = x[DESCRIPTOR_TYPE_UNIFORM_BUFFER-6]
	_ = x[DESCRIPTOR_TYPE_STORAGE_BUFFER-7]
	_ = x[DESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMIC-8]
	_ = x[DESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMIC-9]
	_ = x[DESCRIPTOR_TYPE_INPUT_ATTACHMENT-10]
	_ = x[DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK-1000138000]
	_ = x[DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_KHR-1000150000]
	_ = x[DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_NV-1000165000]
	_ = x[DESCRIPTOR_TYPE_SAMPLE_WEIGHT_IMAGE_QCOM-1000440000]
	_ = x[DESCRIPTOR_TYPE_BLOCK_MATCH_IMAGE_QCOM-1000440001]
	_ = x[DESCRIPTOR_TYPE_MUTABLE_EXT-1000494000]
	_ = x[DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK_EXT-1000138000]
	_ = x[DESCRIPTOR_TYPE_MUTABLE_VALVE-1000494000]
}

const (
	_DescriptorType_name_0 = "DESCRIPTOR_TYPE_SAMPLERDESCRIPTOR_TYPE_COMBINED_IMAGE_SAMPLERDESCRIPTOR_TYPE_SAMPLED_IMAGEDESCRIPTOR_TYPE_STORAGE_IMAGEDESCRIPTOR_TYPE_UNIFORM_TEXEL_BUFFERDESCRIPTOR_TYPE_STORAGE_TEXEL_BUFFERDESCRIPTOR_TYPE_UNIFORM_BUFFERDESCRIPTOR_TYPE_STORAGE_BUFFERDESCRIPTOR_TYPE_UNIFORM_BUFFER_DYNAMICDESCRIPTOR_TYPE_STORAGE_BUFFER_DYNAMICDESCRIPTOR_TYPE_INPUT_ATTACHMENT"
	_DescriptorType_name_1 = "DESCRIPTOR_TYPE_INLINE_UNIFORM_BLOCK"
	_DescriptorType_name_2 = "DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_KHR"
	_DescriptorType_name_3 = "DESCRIPTOR_TYPE_ACCELERATION_STRUCTURE_NV"
	_DescriptorType_name_4 = "DESCRIPTOR_TYPE_SAMPLE_WEIGHT_IMAGE_QCOMDESCRIPTOR_TYPE_BLOCK_MATCH_IMAGE_QCOM"
	_DescriptorType_name_5 = "DESCRIPTOR_TYPE_MUTABLE_EXT"
)

var (
	_DescriptorType_index_0 = [...]uint16{0, 23, 61, 90, 119, 155, 191, 221, 251, 289, 327, 359}
	_DescriptorType_index_4 = [...]uint8{0, 40, 78}
)

func (i DescriptorType) String() string {
	switch {
	case 0 <= i && i <= 10:
		return _DescriptorType_name_0[_DescriptorType_index_0[i]:_DescriptorType_index_0[i+1]]
	case i == 1000138000:
		return _DescriptorType_name_1
	case i == 1000150000:
		return _DescriptorType_name_2
	case i == 1000165000:
		return _DescriptorType_name_3
	case 1000440000 <= i && i <= 1000440001:
		i -= 1000440000
		return _DescriptorType_name_4[_DescriptorType_index_4[i]:_DescriptorType_index_4[i+1]]
	case i == 1000494000:
		return _DescriptorType_name_5
	default:
		return "DescriptorType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET-0]
	_ = x[DESCRIPTOR_UPDATE_TEMPLATE_TYPE_PUSH_DESCRIPTORS_KHR-1000085000]
	_ = x[DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET_KHR-0]
}

const (
	_DescriptorUpdateTemplateType_name_0 = "DESCRIPTOR_UPDATE_TEMPLATE_TYPE_DESCRIPTOR_SET"
	_DescriptorUpdateTemplateType_name_1 = "DESCRIPTOR_UPDATE_TEMPLATE_TYPE_PUSH_DESCRIPTORS_KHR"
)

func (i DescriptorUpdateTemplateType) String() string {
	switch {
	case i == 0:
		return _DescriptorUpdateTemplateType_name_0
	case i == 1000085000:
		return _DescriptorUpdateTemplateType_name_1
	default:
		return "DescriptorUpdateTemplateType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DEVICE_ADDRESS_BINDING_INTERNAL_OBJECT_BIT_EXT-1]
}

const _DeviceAddressBindingFlagBitsEXT_name = "DEVICE_ADDRESS_BINDING_INTERNAL_OBJECT_BIT_EXT"

var _DeviceAddressBindingFlagBitsEXT_index = [...]uint8{0, 46}

func (i DeviceAddressBindingFlagBitsEXT) String() string {
	i -= 1
	if i >= DeviceAddressBindingFlagBitsEXT(len(_DeviceAddressBindingFlagBitsEXT_index)-1) {
		return "DeviceAddressBindingFlagBitsEXT(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _DeviceAddressBindingFlagBitsEXT_name[_DeviceAddressBindingFlagBitsEXT_index[i]:_DeviceAddressBindingFlagBitsEXT_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DEVICE_ADDRESS_BINDING_TYPE_BIND_EXT-0]
	_ = x[DEVICE_ADDRESS_BINDING_TYPE_UNBIND_EXT-1]
}

const _DeviceAddressBindingTypeEXT_name = "DEVICE_ADDRESS_BINDING_TYPE_BIND_EXTDEVICE_ADDRESS_BINDING_TYPE_UNBIND_EXT"

var _DeviceAddressBindingTypeEXT_index = [...]uint8{0, 36, 74}

func (i DeviceAddressBindingTypeEXT) String() string {
	if i < 0 || i >= DeviceAddressBindingTypeEXT(len(_DeviceAddressBindingTypeEXT_index)-1) {
		return "DeviceAddressBindingTypeEXT(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _DeviceAddressBindingTypeEXT_name[_DeviceAddressBindingTypeEXT_index[i]:_DeviceAddressBindingTypeEXT_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DEVICE_DIAGNOSTICS_CONFIG_ENABLE_SHADER_DEBUG_INFO_BIT_NV-1]
	_ = x[DEVICE_DIAGNOSTICS_CONFIG_ENABLE_RESOURCE_TRACKING_BIT_NV-2]
	_ = x[DEVICE_DIAGNOSTICS_CONFIG_ENABLE_AUTOMATIC_CHECKPOINTS_BIT_NV-4]
	_ = x[DEVICE_DIAGNOSTICS_CONFIG_ENABLE_SHADER_ERROR_REPORTING_BIT_NV-8]
}

const (
	_DeviceDiagnosticsConfigFlagBitsNV_name_0 = "DEVICE_DIAGNOSTICS_CONFIG_ENABLE_SHADER_DEBUG_INFO_BIT_NVDEVICE_DIAGNOSTICS_CONFIG_ENABLE_RESOURCE_TRACKING_BIT_NV"
	_DeviceDiagnosticsConfigFlagBitsNV_name_1 = "DEVICE_DIAGNOSTICS_CONFIG_ENABLE_AUTOMATIC_CHECKPOINTS_BIT_NV"
	_DeviceDiagnosticsConfigFlagBitsNV_name_2 = "DEVICE_DIAGNOSTICS_CONFIG_ENABLE_SHADER_ERROR_REPORTING_BIT_NV"
)

var (
	_DeviceDiagnosticsConfigFlagBitsNV_index_0 = [...]uint8{0, 57, 114}
)

func (i DeviceDiagnosticsConfigFlagBitsNV) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _DeviceDiagnosticsConfigFlagBitsNV_name_0[_DeviceDiagnosticsConfigFlagBitsNV_index_0[i]:_DeviceDiagnosticsConfigFlagBitsNV_index_0[i+1]]
	case i == 4:
		return _DeviceDiagnosticsConfigFlagBitsNV_name_1
	case i == 8:
		return _DeviceDiagnosticsConfigFlagBitsNV_name_2
	default:
		return "DeviceDiagnosticsConfigFlagBitsNV(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DEVICE_EVENT_TYPE_DISPLAY_HOTPLUG_EXT-0]
}

const _DeviceEventTypeEXT_name = "DEVICE_EVENT_TYPE_DISPLAY_HOTPLUG_EXT"

var _DeviceEventTypeEXT_index = [...]uint8{0, 37}

func (i DeviceEventTypeEXT) String() string {
	if i < 0 || i >= DeviceEventTypeEXT(len(_DeviceEventTypeEXT_index)-1) {
		return "DeviceEventTypeEXT(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _DeviceEventTypeEXT_name[_DeviceEventTypeEXT_index[i]:_DeviceEventTypeEXT_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DEVICE_FAULT_ADDRESS_TYPE_NONE_EXT-0]
	_ = x[DEVICE_FAULT_ADDRESS_TYPE_READ_INVALID_EXT-1]
	_ = x[DEVICE_FAULT_ADDRESS_TYPE_WRITE_INVALID_EXT-2]
	_ = x[DEVICE_FAULT_ADDRESS_TYPE_EXECUTE_INVALID_EXT-3]
	_ = x[DEVICE_FAULT_ADDRESS_TYPE_INSTRUCTION_POINTER_UNKNOWN_EXT-4]
	_ = x[DEVICE_FAULT_ADDRESS_TYPE_INSTRUCTION_POINTER_INVALID_EXT-5]
	_ = x[DEVICE_FAULT_ADDRESS_TYPE_INSTRUCTION_POINTER_FAULT_EXT-6]
}

const _DeviceFaultAddressTypeEXT_name = "DEVICE_FAULT_ADDRESS_TYPE_NONE_EXTDEVICE_FAULT_ADDRESS_TYPE_READ_INVALID_EXTDEVICE_FAULT_ADDRESS_TYPE_WRITE_INVALID_EXTDEVICE_FAULT_ADDRESS_TYPE_EXECUTE_INVALID_EXTDEVICE_FAULT_ADDRESS_TYPE_INSTRUCTION_POINTER_UNKNOWN_EXTDEVICE_FAULT_ADDRESS_TYPE_INSTRUCTION_POINTER_INVALID_EXTDEVICE_FAULT_ADDRESS_TYPE_INSTRUCTION_POINTER_FAULT_EXT"

var _DeviceFaultAddressTypeEXT_index = [...]uint16{0, 34, 76, 119, 164, 221, 278, 333}

func (i DeviceFaultAddressTypeEXT) String() string {
	if i < 0 || i >= DeviceFaultAddressTypeEXT(len(_DeviceFaultAddressTypeEXT_index)-1) {
		return "DeviceFaultAddressTypeEXT(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _DeviceFaultAddressTypeEXT_name[_DeviceFaultAddressTypeEXT_index[i]:_DeviceFaultAddressTypeEXT_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DEVICE_FAULT_VENDOR_BINARY_HEADER_VERSION_ONE_EXT-1]
}

const _DeviceFaultVendorBinaryHeaderVersionEXT_name = "DEVICE_FAULT_VENDOR_BINARY_HEADER_VERSION_ONE_EXT"

var _DeviceFaultVendorBinaryHeaderVersionEXT_index = [...]uint8{0, 49}

func (i DeviceFaultVendorBinaryHeaderVersionEXT) String() string {
	i -= 1
	if i < 0 || i >= DeviceFaultVendorBinaryHeaderVersionEXT(len(_DeviceFaultVendorBinaryHeaderVersionEXT_index)-1) {
		return "DeviceFaultVendorBinaryHeaderVersionEXT(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _DeviceFaultVendorBinaryHeaderVersionEXT_name[_DeviceFaultVendorBinaryHeaderVersionEXT_index[i]:_DeviceFaultVendorBinaryHeaderVersionEXT_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DEVICE_GROUP_PRESENT_MODE_LOCAL_BIT_KHR-1]
	_ = x[DEVICE_GROUP_PRESENT_MODE_REMOTE_BIT_KHR-2]
	_ = x[DEVICE_GROUP_PRESENT_MODE_SUM_BIT_KHR-4]
	_ = x[DEVICE_GROUP_PRESENT_MODE_LOCAL_MULTI_DEVICE_BIT_KHR-8]
}

const (
	_DeviceGroupPresentModeFlagBitsKHR_name_0 = "DEVICE_GROUP_PRESENT_MODE_LOCAL_BIT_KHRDEVICE_GROUP_PRESENT_MODE_REMOTE_BIT_KHR"
	_DeviceGroupPresentModeFlagBitsKHR_name_1 = "DEVICE_GROUP_PRESENT_MODE_SUM_BIT_KHR"
	_DeviceGroupPresentModeFlagBitsKHR_name_2 = "DEVICE_GROUP_PRESENT_MODE_LOCAL_MULTI_DEVICE_BIT_KHR"
)

var (
	_DeviceGroupPresentModeFlagBitsKHR_index_0 = [...]uint8{0, 39, 79}
)

func (i DeviceGroupPresentModeFlagBitsKHR) String() string {
	switch {
	case 1 <= i && i <= 2:
		i -= 1
		return _DeviceGroupPresentModeFlagBitsKHR_name_0[_DeviceGroupPresentModeFlagBitsKHR_index_0[i]:_DeviceGroupPresentModeFlagBitsKHR_index_0[i+1]]
	case i == 4:
		return _DeviceGroupPresentModeFlagBitsKHR_name_1
	case i == 8:
		return _DeviceGroupPresentModeFlagBitsKHR_name_2
	default:
		return "DeviceGroupPresentModeFlagBitsKHR(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DEVICE_MEMORY_REPORT_EVENT_TYPE_ALLOCATE_EXT-0]
	_ = x[DEVICE_MEMORY_REPORT_EVENT_TYPE_FREE_EXT-1]
	_ = x[DEVICE_MEMORY_REPORT_EVENT_TYPE_IMPORT_EXT-2]
	_ = x[DEVICE_MEMORY_REPORT_EVENT_TYPE_UNIMPORT_EXT-3]
	_ = x[DEVICE_MEMORY_REPORT_EVENT_TYPE_ALLOCATION_FAILED_EXT-4]
}

const _DeviceMemoryReportEventTypeEXT_name = "DEVICE_MEMORY_REPORT_EVENT_TYPE_ALLOCATE_EXTDEVICE_MEMORY_REPORT_EVENT_TYPE_FREE_EXTDEVICE_MEMORY_REPORT_EVENT_TYPE_IMPORT_EXTDEVICE_MEMORY_REPORT_EVENT_TYPE_UNIMPORT_EXTDEVICE_MEMORY_REPORT_EVENT_TYPE_ALLOCATION_FAILED_EXT"

var _DeviceMemoryReportEventTypeEXT_index = [...]uint8{0, 44, 84, 126, 170, 223}

func (i DeviceMemoryReportEventTypeEXT) String() string {
	if i < 0 || i >= DeviceMemoryReportEventTypeEXT(len(_DeviceMemoryReportEventTypeEXT_index)-1) {
		return "DeviceMemoryReportEventTypeEXT(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _DeviceMemoryReportEventTypeEXT_name[_DeviceMemoryReportEventTypeEXT_index[i]:_DeviceMemoryReportEventTypeEXT_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DEVICE_QUEUE_CREATE_PROTECTED_BIT-1]
}

const _DeviceQueueCreateFlagBits_name = "DEVICE_QUEUE_CREATE_PROTECTED_BIT"

var _DeviceQueueCreateFlagBits_index = [...]uint8{0, 33}

func (i DeviceQueueCreateFlagBits) String() string {
	i -= 1
	if i >= DeviceQueueCreateFlagBits(len(_DeviceQueueCreateFlagBits_index)-1) {
		return "DeviceQueueCreateFlagBits(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _DeviceQueueCreateFlagBits_name[_DeviceQueueCreateFlagBits_index[i]:_DeviceQueueCreateFlagBits_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DIRECT_DRIVER_LOADING_MODE_EXCLUSIVE_LUNARG-0]
	_ = x[DIRECT_DRIVER_LOADING_MODE_INCLUSIVE_LUNARG-1]
}

const _DirectDriverLoadingModeLUNARG_name = "DIRECT_DRIVER_LOADING_MODE_EXCLUSIVE_LUNARGDIRECT_DRIVER_LOADING_MODE_INCLUSIVE_LUNARG"

var _DirectDriverLoadingModeLUNARG_index = [...]uint8{0, 43, 86}

func (i DirectDriverLoadingModeLUNARG) String() string {
	if i < 0 || i >= DirectDriverLoadingModeLUNARG(len(_DirectDriverLoadingModeLUNARG_index)-1) {
		return "DirectDriverLoadingModeLUNARG(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _DirectDriverLoadingModeLUNARG_name[_DirectDriverLoadingModeLUNARG_index[i]:_DirectDriverLoadingModeLUNARG_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[DISCARD_RECTANGLE_MODE_INCLUSIVE_EXT-0]
	_ = x[DISCARD_RECTANGLE_MODE_EXCLUSIVE_EXT-1]
}

const _DiscardRectangleModeEXT_name = "DISCARD_RECTANGLE_MODE_INCLUSIVE_EXTDISCARD_RECTANGLE_MODE_EXCLUSIVE_EXT"

var _DiscardRectangleModeEXT_index = [...]uint8{0, 36, 72}

func (i DiscardRectangleModeEXT) String() string {
	if i < 0 || i >= DiscardRectangleModeEXT(len(_DiscardRectangleModeEXT_index)-1) {
		return "DiscardRectangleModeEXT(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _DiscardRectangleModeEXT_name[_DiscardRectangleModeEXT_index[i]:_DiscardRectangleModeEXT_index[i+1]]
}