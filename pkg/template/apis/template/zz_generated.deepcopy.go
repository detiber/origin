// +build !ignore_autogenerated_openshift

// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package template

import (
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	api "k8s.io/kubernetes/pkg/api"
	reflect "reflect"
)

func init() {
	SchemeBuilder.Register(RegisterDeepCopies)
}

// RegisterDeepCopies adds deep-copy functions to the given scheme. Public
// to allow building arbitrary schemes.
func RegisterDeepCopies(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedDeepCopyFuncs(
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_template_BrokerTemplateInstance, InType: reflect.TypeOf(&BrokerTemplateInstance{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_template_BrokerTemplateInstanceList, InType: reflect.TypeOf(&BrokerTemplateInstanceList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_template_BrokerTemplateInstanceSpec, InType: reflect.TypeOf(&BrokerTemplateInstanceSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_template_Parameter, InType: reflect.TypeOf(&Parameter{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_template_Template, InType: reflect.TypeOf(&Template{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_template_TemplateInstance, InType: reflect.TypeOf(&TemplateInstance{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_template_TemplateInstanceCondition, InType: reflect.TypeOf(&TemplateInstanceCondition{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_template_TemplateInstanceList, InType: reflect.TypeOf(&TemplateInstanceList{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_template_TemplateInstanceRequester, InType: reflect.TypeOf(&TemplateInstanceRequester{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_template_TemplateInstanceSpec, InType: reflect.TypeOf(&TemplateInstanceSpec{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_template_TemplateInstanceStatus, InType: reflect.TypeOf(&TemplateInstanceStatus{})},
		conversion.GeneratedDeepCopyFunc{Fn: DeepCopy_template_TemplateList, InType: reflect.TypeOf(&TemplateList{})},
	)
}

// DeepCopy_template_BrokerTemplateInstance is an autogenerated deepcopy function.
func DeepCopy_template_BrokerTemplateInstance(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BrokerTemplateInstance)
		out := out.(*BrokerTemplateInstance)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_template_BrokerTemplateInstanceSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_template_BrokerTemplateInstanceList is an autogenerated deepcopy function.
func DeepCopy_template_BrokerTemplateInstanceList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BrokerTemplateInstanceList)
		out := out.(*BrokerTemplateInstanceList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]BrokerTemplateInstance, len(*in))
			for i := range *in {
				if err := DeepCopy_template_BrokerTemplateInstance(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_template_BrokerTemplateInstanceSpec is an autogenerated deepcopy function.
func DeepCopy_template_BrokerTemplateInstanceSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*BrokerTemplateInstanceSpec)
		out := out.(*BrokerTemplateInstanceSpec)
		*out = *in
		if in.BindingIDs != nil {
			in, out := &in.BindingIDs, &out.BindingIDs
			*out = make([]string, len(*in))
			copy(*out, *in)
		}
		return nil
	}
}

// DeepCopy_template_Parameter is an autogenerated deepcopy function.
func DeepCopy_template_Parameter(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Parameter)
		out := out.(*Parameter)
		*out = *in
		return nil
	}
}

// DeepCopy_template_Template is an autogenerated deepcopy function.
func DeepCopy_template_Template(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*Template)
		out := out.(*Template)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if in.Parameters != nil {
			in, out := &in.Parameters, &out.Parameters
			*out = make([]Parameter, len(*in))
			copy(*out, *in)
		}
		if in.Objects != nil {
			in, out := &in.Objects, &out.Objects
			*out = make([]runtime.Object, len(*in))
			for i := range *in {
				if newVal, err := c.DeepCopy(&(*in)[i]); err != nil {
					return err
				} else {
					(*out)[i] = *newVal.(*runtime.Object)
				}
			}
		}
		if in.ObjectLabels != nil {
			in, out := &in.ObjectLabels, &out.ObjectLabels
			*out = make(map[string]string)
			for key, val := range *in {
				(*out)[key] = val
			}
		}
		return nil
	}
}

// DeepCopy_template_TemplateInstance is an autogenerated deepcopy function.
func DeepCopy_template_TemplateInstance(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*TemplateInstance)
		out := out.(*TemplateInstance)
		*out = *in
		if newVal, err := c.DeepCopy(&in.ObjectMeta); err != nil {
			return err
		} else {
			out.ObjectMeta = *newVal.(*v1.ObjectMeta)
		}
		if err := DeepCopy_template_TemplateInstanceSpec(&in.Spec, &out.Spec, c); err != nil {
			return err
		}
		if err := DeepCopy_template_TemplateInstanceStatus(&in.Status, &out.Status, c); err != nil {
			return err
		}
		return nil
	}
}

// DeepCopy_template_TemplateInstanceCondition is an autogenerated deepcopy function.
func DeepCopy_template_TemplateInstanceCondition(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*TemplateInstanceCondition)
		out := out.(*TemplateInstanceCondition)
		*out = *in
		out.LastTransitionTime = in.LastTransitionTime.DeepCopy()
		return nil
	}
}

// DeepCopy_template_TemplateInstanceList is an autogenerated deepcopy function.
func DeepCopy_template_TemplateInstanceList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*TemplateInstanceList)
		out := out.(*TemplateInstanceList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]TemplateInstance, len(*in))
			for i := range *in {
				if err := DeepCopy_template_TemplateInstance(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_template_TemplateInstanceRequester is an autogenerated deepcopy function.
func DeepCopy_template_TemplateInstanceRequester(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*TemplateInstanceRequester)
		out := out.(*TemplateInstanceRequester)
		*out = *in
		return nil
	}
}

// DeepCopy_template_TemplateInstanceSpec is an autogenerated deepcopy function.
func DeepCopy_template_TemplateInstanceSpec(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*TemplateInstanceSpec)
		out := out.(*TemplateInstanceSpec)
		*out = *in
		if err := DeepCopy_template_Template(&in.Template, &out.Template, c); err != nil {
			return err
		}
		if in.Secret != nil {
			in, out := &in.Secret, &out.Secret
			*out = new(api.LocalObjectReference)
			**out = **in
		}
		if in.Requester != nil {
			in, out := &in.Requester, &out.Requester
			*out = new(TemplateInstanceRequester)
			**out = **in
		}
		return nil
	}
}

// DeepCopy_template_TemplateInstanceStatus is an autogenerated deepcopy function.
func DeepCopy_template_TemplateInstanceStatus(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*TemplateInstanceStatus)
		out := out.(*TemplateInstanceStatus)
		*out = *in
		if in.Conditions != nil {
			in, out := &in.Conditions, &out.Conditions
			*out = make([]TemplateInstanceCondition, len(*in))
			for i := range *in {
				if err := DeepCopy_template_TemplateInstanceCondition(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}

// DeepCopy_template_TemplateList is an autogenerated deepcopy function.
func DeepCopy_template_TemplateList(in interface{}, out interface{}, c *conversion.Cloner) error {
	{
		in := in.(*TemplateList)
		out := out.(*TemplateList)
		*out = *in
		if in.Items != nil {
			in, out := &in.Items, &out.Items
			*out = make([]Template, len(*in))
			for i := range *in {
				if err := DeepCopy_template_Template(&(*in)[i], &(*out)[i], c); err != nil {
					return err
				}
			}
		}
		return nil
	}
}
